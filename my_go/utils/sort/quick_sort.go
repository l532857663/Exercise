package sort

/* 快速排序

时间复杂度：O(nlogn)
空间复杂度：O(logn)?
非稳定排序、原地排序
*/

func quickSort(arr []int, left, right int) {
	// 数组长度小于2，返回
	arrLength := len(arr)
	if arrLength == 0 || arrLength < 2 {
		return
	}
	// 获取指标数据
	flag := arr[left]
	for left < right {
		for left < right && arr[right] >= flag {
			right -= 1
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= flag {
			left += 1
		}
		arr[right] = arr[left]
	}
	arr[left] = flag
	lArr := arr[:left]
	rArr := arr[left+1:]
	quickSort(lArr, 0, len(lArr)-1)
	quickSort(rArr, 0, len(rArr)-1)
	return
}

func QuickSort(arr []int) {
	// 数组长度小于2，返回
	arrLength := len(arr)
	if arrLength == 0 || arrLength < 2 {
		return
	}

	quickSort(arr, 0, len(arr)-1)
	return
}
