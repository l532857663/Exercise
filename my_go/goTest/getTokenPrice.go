package main

import (
	"encoding/json"
	"fmt"
	"os"

	coingecko "github.com/superoo7/go-gecko/v3"
	geckoTypes "github.com/superoo7/go-gecko/v3/types"
)

type CGClient struct {
	HttpCli *coingecko.Client
	BaseURL string
}

type CoinBase []coinBaseStruct
type coinBaseStruct struct {
	ID       string            `json:"id"`
	Symbol   string            `json:"symbol"`
	Name     string            `json:"name"`
	Platform map[string]string `json:"platforms"`
}

func (cg *CGClient) Ping() {
	ping, err := cg.HttpCli.Ping()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Println("result: %v\n", ping.GeckoSays)
	return
}

func (cg *CGClient) GetPrice(tokenList []string) {
	vc := []string{"usd", "cny"}
	sp, err := cg.HttpCli.SimplePrice(tokenList, vc)
	if err != nil {
		fmt.Printf("GetPrice error: %v\n", err)
		return
	}
	fmt.Printf("wch------- test %+v\n", sp)
	bitcoin := (*sp)["bitcoin"]
	eth := (*sp)["ethereum"]
	fmt.Printf("wch-----  bitcoin %+v, %f, %f\n", bitcoin, bitcoin["usd"], bitcoin["cny"])
	fmt.Printf("wch-----  eth %+v, %f, %f\n", eth, eth["usd"], eth["cny"])
	return
}

func (cg *CGClient) GetCoinsList() {
	/*
		list, err := cg.HttpCli.CoinsList()
		if err != nil {
			fmt.Printf("GetCoinsList error: %v\n", err)
		}
	*/
	url := fmt.Sprintf("%s/coins/list?include_platform=true", cg.BaseURL)
	fmt.Printf("wch-------- market url: %v\n", url)
	resp, err := cg.HttpCli.MakeReq(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// var list *geckoTypes.CoinList
	var list *CoinBase
	err = json.Unmarshal(resp, &list)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Available coins:", len(*list))
	f, err := os.Create("./symbolList.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		i := 0
		for _, coin := range *list {
			i++
			if i%100 == 0 {
				fmt.Printf("wch----- i %v\n", i)
				// break
			}
			resInfo := fmt.Sprintf("%+v", coin)
			// fmt.Printf("wch------ coin %s\n", resInfo)
			f.Write([]byte(resInfo))
			f.Write([]byte("\n"))
		}
	}
	return
}

func (cg *CGClient) GetCoinsMarket() {
	vsCurrency := "cny"
	ids := []string{"bitcoin", "ethereum", "huobi-token", "weth", "basis-bond"}
	perPage := 1
	page := 1
	sparkline := true
	pcp := geckoTypes.PriceChangePercentageObject
	priceChangePercentage := []string{pcp.PCP1h, pcp.PCP24h, pcp.PCP7d, pcp.PCP14d, pcp.PCP30d, pcp.PCP200d, pcp.PCP1y}
	order := geckoTypes.OrderTypeObject.MarketCapDesc

	market, err := cg.HttpCli.CoinsMarket(vsCurrency, ids, order, perPage, page, sparkline, priceChangePercentage)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Total coins: ", len(*market))
	for _, coin := range *market {
		fmt.Printf("wch------- coin: %+v\n", coin)
	}
}

func main() {
	fmt.Printf("wch------- Start\n")
	cg := CGClient{
		HttpCli: coingecko.NewClient(nil),
		BaseURL: "https://api.coingecko.com/api/v3",
	}
	// /ping
	// cg.Ping()

	// /simple/price
	// tokenList := []string{"bitcoin", "ethereum"}
	// cg.GetPrice(tokenList)

	// /coins/list
	// cg.GetCoinsList()

	// /coins/market
	cg.GetCoinsMarket()
	fmt.Printf("wch------- End\n")
}
