package main

import (
	"fmt"
	"strings"
)

func main() {
	var nums = [4]int{1, 2, 3, 4}
	fmt.Printf("Length is: %d", len(nums))
	fmt.Println()
	var n [4]int
	for i, element := range n {
		fmt.Printf("%d=%d ", i, element)
	}
	fmt.Println()
	fmt.Printf("%v\n", n)
	fmt.Printf("%#v\n", n)
	var a1 = [10]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [6]float32{-1, -2}
	fmt.Printf("%#v\n", a2)

	var a5 = [...]int{5, 3, 5, 6, 3}
	fmt.Printf("%d\n", len(a5))

	var a55 = [...]int{5,
		3,
		5,
		6,
		3,
	}
	fmt.Printf("%d\n", len(a55))
	a55[0] = 100
	a55[4] = 100
	fmt.Printf("%#v\n", a55)
	fmt.Println()
	for i, v := range a55 {
		fmt.Printf("Index: %d. Value: %d", i, v)
	}
	fmt.Println()
	fmt.Println(strings.Repeat("#", 20))
	for i := 0; i < len(a55); i++ {
		fmt.Printf("Index: %d. Value: %d", i, a55[i])
	}
	fmt.Println()
	balances := [2][3]int{
		{5, 6, 7},
		[3]int{5, 6, 7},
	}
	fmt.Printf("Length: %d\n", len(balances))
	fmt.Printf("Length: %d\n", len(balances[0]))
	grades := [...]int{
		1:  10,
		0:  4,
		2:  5,
		10: 53,
		55}
	fmt.Printf("%#v\n", grades)
}
