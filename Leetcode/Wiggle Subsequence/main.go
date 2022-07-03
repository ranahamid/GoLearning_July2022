package main

import (
	"fmt"
	"math"
)

func main() {
	var nums = []int{1, 2}
	var v = wiggleMaxLength(nums)
	fmt.Println((v))

}
func wiggleMaxLength(nums []int) int {
	var f = 1
	var d = 1
	for i := 1; i < len(nums); i++ {
		if (nums[i]) > nums[i-1] {
			f = d + 1
		} else if nums[i] < nums[i-1] {
			d = f + 1
		}
	}

	return int(math.Max(float64(f), float64(d)))
}
