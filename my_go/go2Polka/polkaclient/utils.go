package polkaclient

import (
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// 处理交易数据详情
func (n *Node) FindCallNamesForCallID(eventID types.EventID) (types.Text, types.Text, error) {
	// 获取元数据(NOTE: meta的数据结构是随节点变化还是随数据变化)
	meta, err := n.Api.RPC.State.GetMetadataLatest()
	if err != nil {
		fmt.Printf("Get metadata error: %+v\n", err)
		return "", "", err
	}
	switch meta.Version {
	case 14:
		for _, v := range meta.AsMetadataV14.Pallets {
			if uint8(v.Index) != eventID[0] {
				continue
			}
			callType := v.Calls.Type.Int64()
			typ := meta.AsMetadataV14.EfficientLookup[callType]
			if len(typ.Def.Variant.Variants) > 0 {
				for _, vars := range typ.Def.Variant.Variants {
					if uint8(vars.Index) == eventID[1] {
						return v.Name, vars.Name, nil
					}
				}
			}
		}
		return "", "", fmt.Errorf("module index %v out of range", eventID[0])
	default:
		return "", "", fmt.Errorf("unsupported metadata version")
	}
}

// 解析返回结果
func (n *Node) GetResultInfo(res string) {
	bz, err := types.HexDecodeString(res)
	if err != nil {
		panic(err)
	}

	fmt.Printf("wch----- bz: %+v, len: %+v\n", bz, len(bz))
	var aa []byte
	aa = append(aa, bz[0:12]...)
	aa = append(aa, bz[16:]...)
	fmt.Printf("wch----- aa: %+v, len: %+v\n", aa, len(aa))

	data := types.NewStorageDataRaw(aa)

	target := &types.AccountInfo{}
	err = types.DecodeFromBytes(data, target)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wch------ target: %+v\n", target)
	fmt.Printf("wch------ Free string: %+v\n", target.Data.Free.String())
	fmt.Printf("wch------ Free uint64: %+v\n", target.Data.Free.Uint64())
	fmt.Printf("wch------ Free byte: %+v\n", target.Data.Free.Bytes())
	// a, _ := types.BigIntToUintBytes(target.Data.Free, 16)
	// fmt.Printf("wch------ a: %+v\n", a)
	// fmt.Println("wch------ a: ", a)
}
