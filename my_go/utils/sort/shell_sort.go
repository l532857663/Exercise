package sort

/* shell sort 希尔排序
希尔排序是插入排序的改进版，实现简单，对于中等规模数据的性能表现还不错
时间复杂度：O(nlogn)
空间复杂度：O(1)
非稳定排序、原地排序
*/

func ShellSort(arr []int) {
	// 数组长度小于2，返回
	arrLength := len(arr)
	if arrLength == 0 || arrLength < 2 {
		return
	}
	// 对数组进行分组处理
	l := arrLength / 2
	for i := l; i > 0; i /= 2 {
		for j := i; j < arrLength; j++ {
			ChangeIndex(arr, i, j)
		}
	}
	return
}

func ChangeIndex(arr []int, h, i int) {
	tmp := arr[i]
	var j int
	for j = i - h; j >= 0 && tmp < arr[j]; j -= h {
		arr[j+h] = arr[j]
	}
	arr[j+h] = tmp

	return
}
