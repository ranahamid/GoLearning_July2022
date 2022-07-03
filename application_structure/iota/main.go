package main

import "fmt"

func main() {
	const (
		c1 = iota
		c2
		c3
	)
	fmt.Println(c1, c2, c3)
	const (
		North = (iota * 2) + 1
		_
		East
		South
		West
	)
	fmt.Print(West)
}
