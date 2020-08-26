package main

import "fmt"

func selection(nums []int) []int {
	numLen := len(nums)

	for i := 0; i < numLen; i++ {

		var min int //每次遍历的最小值
		var k int   //每次最小值的位置index

		for j := i; j < numLen; j++ {
			if j == i || nums[j] < min {
				min = nums[j]
				k = j
			}
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	return nums
}

func main() {
	num := []int{10, 18, 7, 8, 5, 1, 3, 6, 4}
	res := selection(num)
	fmt.Println(res)
}
