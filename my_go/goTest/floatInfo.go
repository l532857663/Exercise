package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func StringToBigRat(data string) *big.Rat {
	var (
		res = new(big.Rat)
		ok  bool
	)
	res, ok = res.SetString(data)
	if !ok {
		return res
	}
	return res
}

func CalcBalanceByDecimal(amount string, decimal, reserverd int) string {
	// 获取小数位对应值10^decimal的有理数
	d := new(big.Rat).SetFloat64(math.Pow10(decimal))
	// 获取余额的有理数
	a := StringToBigRat(amount)
	// 计算余额的对应小数位值
	res := new(big.Rat).Quo(a, d)
	// 返回字符串结果
	var resStr string
	if reserverd > 0 {
		resStr = res.FloatString(reserverd)
	} else {
		resFloat, _ := res.Float64()
		fmt.Println("resFloat:", resFloat)
		resStr = strconv.FormatFloat(resFloat, 'f', -1, 64)
	}
	return resStr
}

// 小数位转化
func Float2String() {
	a := "27300000000000"
	d := 18
	res := CalcBalanceByDecimal(a, d, 0)
	fmt.Println("res:", res)
	return
}

func main() {
	Float2String()
}
