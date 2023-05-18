package goBTC

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// 解析Input Script inscribe
func GetScriptString(data string) []string {
	char, err := hex.DecodeString(data)
	if err != nil {
		fmt.Println("The string not hex string")
		return nil
	}
	var (
		resList  = []string{}
		res      = ""
		i        = 0
		headChar = char[0]
		dataType string //铭文类型字符串
	)
	// 初步认定格式
	if headChar == 0x20 {
		// 头部数据(推测是地址)
		res, i = getHexData(char, i, "B")
		resList = append(resList, res)
	} else {
		fmt.Printf("HeadChar: %x\n", headChar)
		return resList
	}
	// 中段数据
	flage1 := hex.EncodeToString(char[i : i+2])
	if flage1 == "ac00" {
		i += 3 // 跳过ac0063
	} else if flage1 == "ac06" {
		i += 11 // 跳过ac06****0063
	} else {
		fmt.Printf("flage1: %s\n", flage1)
		return resList
	}
	// Get 铭文 ord
	res, i = getHexData(char, i, "S")
	resList = append(resList, res)
	// Get 铭文类型
	i += 2 // 跳过 0101
	dataType, i = getHexData(char, i, "S")
	resList = append(resList, dataType)
	tList := strings.Split(dataType, "/")
	inscribeType := tList[0]
	IType, ok := InscribeTypeMap[inscribeType]
	if res != "ord" || !ok {
		fmt.Printf("%s: %s\n", res, dataType)
		return resList
	}

	// 后段数据
	if char[i] == 0x00 {
		i++ // OP_FALSE
	}
	lenChar := len(char[i:])
	res = ""
	start := 0
	fmt.Printf("filer lenChar: %+v, char[i]: %x, i: %+v\n", lenChar, char[i], i)
	for j := 0; j <= lenChar; j += len(char[start:i]) {
		start = i
		tmp := ""
		if char[i] >= 0x01 && char[i] <= 0x4b {
			tmp, i = getHexData(char, i, "S")
		} else if char[i] == 0x4c {
			i++ // OP_PUSHDATA1
			tmp, i = getHexData(char, i, "S")
		} else if char[i] == 0x4d {
			hexLen := getHexLen(char[i+1 : i+3])
			i += 3 // OP_PUSHDATA2
			tmp, i = getHexPushData(char, hexLen, i, IType)
		} else if char[i] == 0x4e {
			hexLen := getHexLen(char[i+1 : i+5])
			i += 5 // OP_PUSHDATA4
			tmp, i = getHexPushData(char, hexLen, i, IType)
		} else {
			opCode := char[i]
			if opCode == 0x68 {
				fmt.Println("End!")
				break
			}
			fmt.Printf("End?: %x, %s, [%s]\n", opCode, string(opCode), BTCScriptValueMap[opCode])
			break
		}
		res += tmp
	}
	// 获取铭文数据
	resList = append(resList, res)
	return resList
}

// 获取数据
func getHexData(char []byte, i int, getType string) (string, int) {
	hexLen := int(char[i])
	i++
	return getHexPushData(char, hexLen, i, getType)
}

func getHexPushData(char []byte, hexLen, i int, getType string) (string, int) {
	data := char[i : i+hexLen]
	i += hexLen
	res := ""
	switch getType {
	case "B":
		res = hex.EncodeToString(data)
	case "S":
		res = string(data)
	}
	return res, i
}

func getHexLen(char []byte) int {
	var num uint32
	for i := len(char) - 1; i >= 0; i-- {
		num = num<<8 + uint32(char[i]) // 将字节数组转换为数字
	}
	return int(num)
}

func GetWitnessScript(data string) []string {
	char, err := hex.DecodeString(data)
	if err != nil {
		fmt.Println("The string not hex string")
		return nil
	}
	for _, c := range char {
		fmt.Printf("char: %x   %s   %s\n", c, string(c), BTCScriptValueMap[c])
	}
	return nil
}
