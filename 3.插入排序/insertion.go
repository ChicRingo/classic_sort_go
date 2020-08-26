package main

import "fmt"

func insertion(nums []int) []int {

	// 自己写法
	numLen := len(nums)
	for i := 1; i < numLen; i++ { //从第2个开始，i=1开始

		temp := nums[i] // 定义一个临时变量取当前的值
		// 把该值tmp与之前的数 从后往前进行比较
		for j := i - 1; j >= 0; j-- {
			// 当 当前值tmp 比之前的某个值大,那么插入,否则把每个值往后移一位
			if temp > nums[j] {
				nums[j+1] = temp
				break
			} else {
				nums[j+1] = nums[j]
			}
			// 如果都移动了,然后遍历到0了,说明当前值是最小的,把当前值放到最小的位置
			if j == 0 {
				nums[0] = temp
			}
		}
	}

	// 网上写法
	//for i := range nums {
	//	preIndex := i - 1  // 前一个索引值
	//	current := nums[i] // 定义一个临时变量取当前的值
	//	for preIndex >= 0 && nums[preIndex] > current {
	//		nums[preIndex+1] = nums[preIndex]
	//		preIndex -= 1
	//	}
	//	nums[preIndex+1] = current
	//}

	return nums
}

func main() {
	num := []int{10, 18, 7, 8, 5, 1, 3, 6, 4}
	res := insertion(num)
	fmt.Println(res)
}
