package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	var ls int
	for i := range nums {
		ls = m[nums[i]]
		if ls != 0 {
			return []int{ls - 1, i}
		}
		m[target-nums[i]] = i + 1
	}
	return nil
}

func main() {
	// define parameter
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Printf("Params:\n  nums: %v\n  target: %v\n", nums, target)

	result := twoSum(nums, target)

	fmt.Printf("Result:\n  %v\n", result)
}
