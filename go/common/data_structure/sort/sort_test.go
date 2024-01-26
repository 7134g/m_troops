package sort

import (
	"fmt"
	"testing"
	"time"
)

func TestMer(t *testing.T) {
	// 归并
	arr := []int{3, 9, 1, 4, 7, 6, 8, 2, 5}
	fmt.Println("Before sorting:", arr)
	arr = mergeSort(arr)
	fmt.Println("After sorting:", arr)
}

func TestBub(t *testing.T) {
	// 冒泡
	arr := []int{5, 3, 8, 4, 2}
	bubblingSortDown(arr)
	fmt.Println(arr) // 输出 [2 3 4 5 8]
}

func TestQui(t *testing.T) {
	// 快排
	arr := []int{5, 3, 8, 4, 2}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr) // 输出 [2 3 4 5 8]
}

func TestIns(t *testing.T) {
	// 插排
	var arr = []int{19, 13, 27, 15, 3, 4, 26, 12, 1, 0}
	fmt.Println(time.Now().String(), "排序前：", arr)
	fmt.Println(time.Now().String(), "排序后:", insertSort(arr))
}
