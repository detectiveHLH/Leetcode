package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	length := len(nums)
	l := 0

	for i := 1; i < length; i++ {
		if nums[i] != nums[l] {
			l++
			nums[l] = nums[i]
		}
	}
	return l + 1;
}

func main() {
	// define parameter
	num := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Printf("Params:\n  nums: %v\n", num)

	index := removeDuplicates(num)

	fmt.Printf("Result:\n ")
	for i := 0; i < index; i++ {
		fmt.Print(" ", num[i], " ")
	}
	fmt.Printf("\n")


}
