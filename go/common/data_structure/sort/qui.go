package sort

// 快排
func quickSort(arr []int, left, right int) {
	if left < right {
		// 将数组划分为两个子序列
		pivot := partition(arr, left, right)
		// 递归排序左侧子序列
		quickSort(arr, left, pivot-1)
		// 递归排序右侧子序列
		quickSort(arr, pivot+1, right)
	}
}

// 划分子序列
func partition(arr []int, left, right int) int {
	// 选择最右侧的元素作为基准值
	pivot := arr[right]
	// 初始化 i 和 j 指针
	i := left - 1
	for j := left; j <= right-1; j++ {
		// 如果当前元素小于等于基准值，将其与 i 指针所指向的元素交换位置
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// 将基准值与 i 指针所指向的元素交换位置
	arr[i+1], arr[right] = arr[right], arr[i+1]
	// 返回基准值的下标
	return i + 1
}
