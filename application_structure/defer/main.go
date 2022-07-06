package main

import (
	"fmt"
	"os"
)

func foo() {
	fmt.Println("foo")
}
func bar() {
	fmt.Println("bar")
}
func foobar() {
	fmt.Println("foobar")
}
func foobar1() {
	fmt.Println("foobar1")
}
func main() {
	defer foo()
	bar()
	fmt.Println("main")

	defer foobar()
	defer foobar1()

	file, err := os.Open("main.go")
	if err != nil {

	}
	defer file.Close()

	func(msg string) {
		fmt.Println(msg)
	}("Hello World")
	a := increment(10)
	fmt.Printf("%T\n", a)
	fmt.Printf("%d\n", a())
	fmt.Printf("%d\n", a())
	fmt.Printf("%d\n", a())
	fmt.Printf("%d\n", a())
}

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}
