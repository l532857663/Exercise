package main

import (
	"fmt"
	"goBTC"
	"log"
	"net/http"
)

var srv *goBTC.BTCClient

func init() {
	// 构建节点客户端
	nodeInfo := goBTC.BTC_GETBLOCK_MAIN
	var err error
	srv, err = goBTC.NewBTCClient(nodeInfo)
	if err != nil {
		fmt.Printf("NewBTCClient error: %+v, nodeInfo: %+v\n", err, nodeInfo)
		return
	}
	// 初始化脚本map
	goBTC.InitBtcScriptMap()
}

func GetBlockInfo() {
	// hash := "000000000000000000029730547464f056f8b6e2e0a02eaf69c24389983a04f5"
	// blockInfo, err := srv.GetBlockInfoByHash(hash)
	height := int64(767430)
	blockInfo, err := srv.GetBlockInfoByHeight(height)
	if err != nil {
		fmt.Printf("GetBlockInfoByHash error: %+v\n", err)
		return
	}
	fmt.Printf("%+v\n", blockInfo.Header)
	for i, tx := range blockInfo.Transactions {
		res := goBTC.GetTxWitness(tx)
		if res == "" {
			continue
		}
		resList := goBTC.GetScriptString(res)
		if len(resList) == 4 {
			fmt.Printf("[%d] %s : %s\n", i, resList[1], resList[2])
		}
	}
}

func GetWitnessResByHash(hash string) (string, error) {
	// 查询Witness的铭文数据
	data, err := srv.GetRawTransactionByHash(hash)
	if err != nil {
		fmt.Printf("GetRawTransactionByHash error: %+v\n", err)
		return "", err
	}
	witness := goBTC.GetTxWitnessByTxHex(data.Hex)
	resList := goBTC.GetScriptString(witness)
	fmt.Printf("resList len: %+v\n", len(resList))
	if len(resList) < 4 {
		return "", fmt.Errorf("Not get Inscribe")
	}
	return resList[3], nil
}

func GetInscribeHttp() {
	http.HandleFunc("/webp", webpHandler)

	fmt.Println("Web listen port: 4396")
	log.Fatal(http.ListenAndServe(":4396", nil))
}

func webpHandler(w http.ResponseWriter, r *http.Request) {
	//// // var buf bytes.Buffer
	//// var width, height int
	//// // var data []byte
	//// var err error
	//// hash := "ff4d5e838adfe81c8486ed8630be945badf9a5e75d07262f9d56964eba6ca032" // IMAGE_1
	//// res, err := GetWitnessResByHash(hash)
	//// if err != nil {
	//// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	//// 	return
	//// }
	//// data, err := hex.DecodeString(res)
	//// if err != nil {
	//// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	//// 	return
	//// }

	//// // GetInfo
	//// if width, height, _, err = webp.GetInfo(data); err != nil {
	//// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	//// 	return
	//// }
	//// fmt.Printf("width = %d, height = %d\n", width, height)

	//// // GetMetadata
	//// if metadata, err := webp.GetMetadata(data, "ICCP"); err != nil {
	//// 	fmt.Printf("Metadata: err = %v\n", err)
	//// } else {
	//// 	fmt.Printf("Metadata: %s\n", string(metadata))
	//// }

	//// // Decode webp
	//// img, err := webp.Decode(bytes.NewReader(data))
	//// if err != nil {
	//// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	//// 	return
	//// }
	//// if err := png.Encode(w, img); err != nil {
	//// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	//// 	return
	//// }

	// // Encode lossless webp
	// if err = webp.Encode(&buf, m, &webp.Options{Lossless: true}); err != nil {
	// 	log.Println(err)
	// }
	// if err = ioutil.WriteFile("output.webp", buf.Bytes(), 0666); err != nil {
	// 	log.Println(err)
	// }

	fmt.Println("Save output.webp ok")
}

func GetWitness() {
	// hash := "7fb631b7ed420c07b546ee4db8527a9523bbc44961f9983430166988cd6beeeb" // TEXT_1
	// hash := "bdbf2d7e385f650cbcba9a0ae83dc3f466dadc1e48732835e977cfefe2710b42" // TEXT_2
	// hash := "885441055c7bb5d1c54863e33f5c3a06e5a14cc4749cb61a9b3ff1dbe52a5bbb" // TEXT_3
	// hash := "ff4d5e838adfe81c8486ed8630be945badf9a5e75d07262f9d56964eba6ca032" // IMAGE_1
	// hash := "67df85eb1a66211b4e761d0b76464e5d07e758426214dab5d6fe42b664d979a4" // AUDIO_1
	// hash := "38d89d0506a5c936867b8a8c13b57d815cb2b2d86aee076ffec86b31c2cf51b5" // AUDIO_2
	hash := "98005eed3ff26d93861be4e72f6931795bcb2652d0206bd00230293f1749c9ec"
	GetWitnessResByHash(hash)
}

func main() {
	fmt.Println("vim-go")
	// GetBlockInfo()
	// SignTx()
	GetWitness()
	// GetInscribeHttp()
}
