package common

import (
	"math/rand"
	"time"
)

/**
 * GetRandomString 获取随机字符串 (字符串长度，数据类型[0：纯数字, 1：纯字母，2：数字字母混合])
 */

//
func GetRandomString(length, strType int) string {
	// 回传结果
	result := make([]byte, length)

	isAll := strType < 0 || strType >= 2
	index := [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i += 1 {
		if isAll {
			strType = r.Intn(3)
		}

		scope := index[strType][0]
		base := index[strType][1]

		result[i] = uint8(base + r.Intn(scope))
	}

	return string(result)
}
