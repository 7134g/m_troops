package sort

// 插排
func insertSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
			//fmt.Println(i, j, arr)
		}
	}

	return arr
}
