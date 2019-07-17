package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var arr []int
	input := x
	for x != 0 {
		arr = append(arr, x % 10)
		x /= 10
	}

	var num int
	for index, value := range arr {
		m := len(arr) - index - 1
		num += int(float64(value) * math.Pow(10, float64(m)))
	}

	return input == num
}

// 优化版
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	input := x
	var num int

	for x != 0 {
		n := x % 10
		x /= 10
		num = num * 10 + n
	}

	return num == input
}

func main() {
	// define parameter
	num := 121

	fmt.Printf("Params:\n  nums: %v\n", num)

	result := isPalindrome(num)

	fmt.Printf("Result:\n  %v\n", result)
}

