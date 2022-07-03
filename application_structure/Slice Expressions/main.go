package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	//a[start:stop]
	b := a[1:3]
	fmt.Print(b)
	fmt.Println()
	fmt.Print(a)

	fmt.Println()
	b[0] = 100
	fmt.Print(b)
	fmt.Println()
	fmt.Print(a)

	fmt.Println()
	fmt.Print(a[1:])
	fmt.Println()
	fmt.Print(a[:3])

	fmt.Println()
	fmt.Print(a[:])

	var s1 = append(b[:3], 1030)
	fmt.Println()
	fmt.Print(s1)
}
