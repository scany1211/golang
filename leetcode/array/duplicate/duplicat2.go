package main

import "fmt"

// # 题目描述
// # 找出数组中重复的数字。
// # 在一个长度为 n 的数组 nums 里的所有数字都在 0 ～ n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
// # 示例 1：
// # 输入：
// # [2, 3, 1, 0, 2, 5, 3]
// # 输出：2 或 3
// # 解法
// # 0 ～ n-1 范围内的数，分别还原到对应的位置上，如：数字 2 交换到下标为 2 的位置。
func duplicate(nums []int) int {
	for i, v := range nums {
		for i != v {
			if nums[v] == nums[i] {
				return v
			}
			nums[i], nums[v] = nums[v], nums[i]
			v = nums[i]
			fmt.Println(i, v, nums)
		}
	}
	return -1
}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	duplicate(nums)
}
