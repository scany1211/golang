package main

import "fmt"

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

// 示例:

// 给定 nums = [2, 7, 11, 15], target = 9

// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]
func sum(nums []int, target int) []int {
	m := make(map[int]int)
	var r []int
	fmt.Println("初始m值为：", m)
	for i, _ := range nums {
		diff := target - nums[i]
		if j, ok := m[diff]; ok {
			r = append([]int{i, j}, r...)
			fmt.Println("r的中间状态", r)
		} else {
			m[nums[i]] = i
			fmt.Println("m的中间状态", m)
		}
	}
	return r
}

func main() {
	nums := []int{2, 7, 11, 15, 6, 4}
	target := 17
	fmt.Println(sum(nums, target))
}
