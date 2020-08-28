package main

import "fmt"

// 冒泡排序
func bubble(nums []int) []int {
	numLen := len(nums)

	for i := 0; i < numLen; i++ {
		// 按相邻两个一起比较，所以最后一个不用遍历 -1,已经排序的也不用重新排序，所以-i，不写也可
		for j := 0; j < numLen-1-i; j++ {
			// 当前值大于后一个值，则对换位置
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}

func main() {
	num := []int{10, 8, 7, 18, 5, 1, 3, 6, 4}
	res := bubble(num)
	fmt.Println(res)
}
