package main

import (
	"fmt"
)

func partition(nums []int, left, right int) int {
	privot := left // 设基准值 privot 为left
	index := privot + 1
	for i := index; i <= right; i++ {
		if nums[i] < nums[privot] {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	nums[privot], nums[index-1] = nums[index-1], nums[privot]
	return index - 1
}

func quickSort(nums []int, left, right int) {
	// 如果数组长度小于等于1，说明这个数组分区不需要再排序了，直接return
	if len(nums) <= 1 {
		return
	}
	// 如果左小于右，
	if left < right {
		partIndex := partition(nums, left, right)
		quickSort(nums, left, partIndex-1)
		quickSort(nums, partIndex+1, right)
	}
	return
}

func main() {
	num := []int{10, 18, 7, 8, 5, 1, 3, 6, 4}
	quickSort(num, 0, len(num)-1)
	fmt.Println(num)
}
