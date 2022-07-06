package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("NO ERROR")
}

func arrayPairSum(nums []int) int {

	sort.Ints(nums)
	var sum = 0

	for i := 0; i < len(nums); i = i + 2 {
		sum += nums[i]
	}
	return sum
}

func runningSum(nums []int) []int {
	for i := 0; i < len(nums); i = i + 1 {
		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}

func removeElement(nums []int, val int) int {

	var counter = 0
	for i := 0; i < len(nums); i = i + 1 {
		if nums[i] != val {
			nums[counter] = nums[i]
			counter++
		}
	}
	return counter
}
