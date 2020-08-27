package main

import "fmt"

func selection(nums []int) []int {
	numLen := len(nums)

	for i := 0; i < numLen-1; i++ { //最后一个不用排序了

		var minIndex = i //每次最小值的位置index

		for j := i + 1; j < numLen; j++ { // 从第二个开始排序，也就是i+1
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}

func main() {
	num := []int{10, 18, 7, 8, 5, 1, 3, 6, 4}
	res := selection(num)
	fmt.Println(res)
}
