package main

import "fmt"

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
