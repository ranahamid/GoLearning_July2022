package main

import "fmt"

func main() {
	var nums []int
	fmt.Printf("%#v\n", nums)

	fmt.Printf("Length: %d, capacity: %d\n", len(nums), cap(nums))

	nums = append(nums, 1, 2)
	fmt.Printf("Length: %d, capacity: %d\n", len(nums), cap(nums))

	nums = append(nums, 3)
	fmt.Printf("Length: %d, capacity: %d\n", len(nums), cap(nums))
	nums = append(nums, 53)
	fmt.Printf("Length: %d, capacity: %d\n", len(nums), cap(nums))

	nums = append(nums, 5144)
	fmt.Printf("Length: %d, capacity: %d\n", len(nums), cap(nums))
	nums = append(nums, 5144)
	nums = append(nums, 5144)
	nums = append(nums, 5144)
	fmt.Printf("Length: %d, capacity: %d\n", len(nums), cap(nums))

	nums = append(nums, 5144 44)
	fmt.Printf("Length: %d, capacity: %d\n", len(nums), cap(nums))
}
