package polkaclient

import (
	"fmt"
	"math/big"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/shopspring/decimal"
)

// 通过公钥签名交易
func (n *Node) DoTransfer(priKey, pubKey string) error {
	api := n.Api
	// 创建交易所需的公共参数 -------------------
	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		fmt.Printf("Get metadata error: %+v\n", err)
		return err
	}

	// 使用公钥获取交易to地址结构 toAddr
	toAddr, err := types.NewMultiAddressFromHexAccountID(pubKey)
	if err != nil {
		fmt.Printf("The pubKey is Invalid: %+v\n", err)
		return err
	}
	fmt.Printf("wch---- toAddr: %+v\n", toAddr)
	// 0.1 unit of transfer
	bal, ok := new(big.Int).SetString("20000000000", 10)
	if !ok {
		err = fmt.Errorf("Failed to convert balance")
		return err
	}
	// c, err := types.NewCall(meta, "Balances.transfer", toAddr, types.NewUCompact(bal))
	c, err := types.NewCall(meta, "Balances.transfer_keep_alive", toAddr, types.NewUCompact(bal))
	if err != nil {
		fmt.Printf("New Call error: %+v\n", err)
		return err
	}
	fmt.Printf("wch----- c: %+v\n", c)
	ext := types.NewExtrinsic(c)
	fmt.Printf("wch----- ext: %+v\n", ext)

	// 使用私钥获取地址信息 fromAddr
	fromAddr, accountInfo, err := n.GetAccountKeyPairFromSecret(priKey)
	if err != nil {
		fmt.Printf("Get account key pair from secret error: %+v\n", err)
		return err
	}
	fmt.Printf("wch----- fromaddr: %+v\n, key: %+v\n", fromAddr, types.HexEncodeToString(fromAddr.PublicKey[:]))

	// 获取账户信息
	fmt.Printf("wch----- account: %+v\n", accountInfo)

	rv, err := api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		fmt.Printf("Get runtime version latest error: %+v\n", err)
		return err
	}
	fmt.Printf("wch----- rv: %+v\n", rv)
	genesisHash, err := api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		fmt.Printf("Get block hash by number error: %+v\n", err)
		return err
	}
	fmt.Printf("wch----- genesisHash: %#x\n", genesisHash)

	// 构建交易数据
	o := types.SignatureOptions{
		BlockHash:          genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        genesisHash,
		Nonce:              types.NewUCompactFromUInt(accountInfo.Nonce),
		SpecVersion:        rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: rv.TransactionVersion,
	}
	err = ext.Sign(*fromAddr, o)
	if err != nil {
		fmt.Printf("Get Storage latest by key error: %+v\n", err)
		return err
	}
	fmt.Printf("wch---- END ext: %+v\n", ext)
	hash, err := api.RPC.Author.SubmitExtrinsic(ext)
	if err != nil {
		fmt.Printf("Submit extrinsic error: %+v\n", err)
		return err
	}

	fmt.Printf("Transfer sent with hash %#x\n", hash)
	return nil
}

// 解析block数据
func (n *Node) ShowBlockInfo(theHash string, block *types.SignedBlock) error {
	fmt.Printf("wch------- block: %+v\n", block.Justification)
	header := block.Block.Header
	extrinsics := block.Block.Extrinsics
	// 解析block头数据
	theHeader := &BlockHeader{
		Number:            int64(header.Number),
		BlockHash:         theHash,
		ParentHash:        header.ParentHash.Hex(),
		StateRoot:         header.StateRoot.Hex(),
		ExtrinsicsRoot:    header.ExtrinsicsRoot.Hex(),
		TransactionLength: len(extrinsics),
	}
	fmt.Printf("wch---- the block head info: %+v\n\n", theHeader)
	// 解析block交易数据
	for _, transfer := range extrinsics {
		sinature := transfer.Signature
		// 判断是否是用户交易
		if !sinature.Signer.IsID {
			continue
		}
		fromAddr, err := n.GetAccountInfoByMultiAddress(sinature.Signer)
		if err != nil {
			fmt.Printf("Get account info error: %+v\n", err)
			continue
		}
		fmt.Printf("wch----- fromAddr: %+v\n", fromAddr.Address)
		// 数据
		signature := sinature.Signature
		fmt.Printf("wch---- All transfer: %+v\n", transfer)
		// fmt.Printf("wch---- data: %+v\n", signature)
		if signature.IsEd25519 {
			arg := signature.AsEd25519[:]
			fmt.Printf("wch---- AsEd25519: [%+v]\n", types.HexEncodeToString(arg))
		}
		if signature.IsSr25519 {
			arg := signature.AsSr25519[:]
			fmt.Printf("wch---- AsSr25519: [%+v]\n", arg)
			fmt.Printf("wch---- AsSr25519: [%+v]\n", types.HexEncodeToString(arg))
		}
		// 解析交易类型
		var en types.EventID
		en = [2]byte{
			byte(transfer.Method.CallIndex.SectionIndex),
			byte(transfer.Method.CallIndex.MethodIndex),
		}
		t1, t2, err := n.FindCallNamesForCallID(en)
		if err != nil {
			fmt.Printf("Find event names for event id error: %+v\n", err)
			return err
		}
		transferType := fmt.Sprintf("%v(%v)", t1, t2)
		// 解析交易详情信息
		argsBytes := transfer.Method.Args[:]
		args := &CallArgs{}
		err = types.DecodeFromBytes(argsBytes, args)
		if err != nil {
			fmt.Printf("Decode from bytes for call args error: %+v\n", err)
			return err
		}

		fmt.Printf("wch---- args struct: [%+v]\n", args)
		toAddr, err := n.GetAccountInfoByMultiAddress(args.Dest)
		if err != nil {
			fmt.Printf("Get account info error: %+v\n", err)
			continue
		}
		fmt.Printf("wch----- toAddr: %+v\n", toAddr.Address)
		value := decimal.New(args.Value.Int64(), -10)

		theTransfer := &BlockTransaction{
			Version:      transfer.Version,
			Nonce:        uint64(sinature.Nonce.Int64()),
			TransferType: transferType,
			FromAddress:  fromAddr.Address,
			ToAddress:    toAddr.Address,
			Value:        value.String(),
		}
		fmt.Printf("wch---- the transfer: %+v\n", theTransfer)
	}

	return nil
}
