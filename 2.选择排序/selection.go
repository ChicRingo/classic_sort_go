package main

import "fmt"

func selection(nums []int) []int {
	numLen := len(nums)
	//最后一个不用排序了
	for i := 0; i < numLen-1; i++ {
		//每次最小值的位置index
		var minIndex = i
		// 从第二个开始排序，也就是i+1
		for j := i + 1; j < numLen; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}

	return nums
}

/*
表现最稳定的排序算法之一，因为无论什么数据进去都是O(n2)的时间复杂度，
所以用到它的时候，数据规模越小越好。
唯一的好处可能就是不占用额外的内存空间了吧。
理论上讲，选择排序可能也是平时排序一般人想到的最多的排序方法了吧。
*/

func main() {
	num := []int{10, 18, 7, 8, 5, 1, 3, 6, 4}
	res := selection(num)
	fmt.Println(res)
}
