package main

import (
	"fmt"
)

func main() {
	const a float64 = 5.1
	fmt.Println(a)
	const b = 6
	fmt.Println(b)
	const c = a * b
	fmt.Printf("%T", c)

	fmt.Println(c)
	const st = "Hello " + "Go"
	fmt.Println(st)
	const d = 5 > 10
	fmt.Println(d)

	var i int = 5
	var j float64 = float64(i)
	fmt.Printf("%d %f", i, j)
}
