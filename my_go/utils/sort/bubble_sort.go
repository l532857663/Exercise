package sort

/* 冒泡排序

时间复杂度：O(n^2)
空间复杂度：O(1)
稳定排序、原地排序
*/

func BubbleSort(arr []int) {
	// 数组长度小于2，返回
	arrLength := len(arr)
	if arrLength == 0 || arrLength < 2 {
		return
	}
	// 判断是否发生交换
	var flag bool
	// 循环判断
	for i := arrLength; i > 0 && !flag; i-- {
		flag = true
		for j := 0; j < i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
		}
	}
	return
}
