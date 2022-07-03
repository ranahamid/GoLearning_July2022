package main

import (
	"fmt"
)

func main() {
	var i int8 = 127
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	k := 5
	fmt.Println(k)
	var num = []int{1, 3}
	_ = num
	var m = map[string]int{
		"1": 5,
	}
	_ = m
	//ARRAY
	var nums = []int{1, 2, 3, 4}
	fmt.Printf("%T\n", nums)
	var stChars = []string{"a", "c"}
	fmt.Printf("%T\n", stChars)
	//map
	var balances = map[string]float64{
		"USD": 2,
	}
	fmt.Printf("%T\n", balances)

	var dictionary = map[string]int{
		"a": 1,
	}
	fmt.Printf("%T\n", dictionary)
	var dic = map[int]int{
		1: 10,
		2: 20,
	}
	fmt.Printf("%T\n", dic)

	//struct
	type person struct {
		name string
		age  int
	}
	var you person
	you.age = 10
	fmt.Printf("%T\n", you)
	fmt.Printf("%d\n", you.age)
	//pointer
	var x int = 2
	var pt = &x
	fmt.Printf("%T %v\n", pt, pt)
	fmt.Printf("%T\n", f)
}
func f() {

}
