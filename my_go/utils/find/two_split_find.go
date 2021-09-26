package find

/* 二分查找
注：首选是个有序数组,从小到大
*/

func TwoSplitFInd(arr []int, find int) (int, bool) {
	arrLen := len(arr)
	if arrLen < 1 {
		return -1, false
	}
	// 判断中间值
	flag := arrLen / 2
	if flag <= 1 {
		if arr[0] == find {
			return 0, true
		} else if arr[1] == find {
			return 1, true
		} else {
			return -1, false
		}
	}
	// 数值大，往左找。数值小，往右找
	if arr[flag] > find {
		index, ok := TwoSplitFInd(arr[:flag], find)
		return index, ok
	} else if arr[flag] < find {
		index, ok := TwoSplitFInd(arr[flag+1:], find)
		return flag + 1 + index, ok
	} else {
		return flag, true
	}
	return -1, false
}
