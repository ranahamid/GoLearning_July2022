package main

import "fmt"

func main() {
	const day int = 9
	fmt.Println(day)

	const pi float64 = 3.141592654
	const x = 5
	const y = 0.3
	fmt.Println(x / y)

	const (
		c1 = 8
		c2
		c3
	)
	fmt.Println(c1, c2, c3)
	//c2=39

}
