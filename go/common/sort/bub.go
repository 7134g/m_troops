package sort

import "fmt"

// 冒泡
// 从大到小
func bubblingSortUP(tar []int) []int {
	for i := 0; i < len(tar); i++ {
		l := len(tar) - i - 1
		for j := 0; j < l; j++ {
			if tar[j] < tar[j+1] {
				tar[j], tar[j+1] = tar[j+1], tar[j]
			}
			fmt.Println(i, l, tar)
		}
	}
	return tar
}

// 从小到大
func bubblingSortDown(tar []int) []int {
	for i := 0; i < len(tar); i++ {
		for j := 0; j < len(tar)-i-1; j++ {
			if tar[j] > tar[j+1] {
				tar[j], tar[j+1] = tar[j+1], tar[j]
			}
			//fmt.Println(tar)
		}
	}
	return tar
}
