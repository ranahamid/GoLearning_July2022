package main

import (
	"fmt"
)

func main() {
	var cities []string
	fmt.Println("cities is equal to nil", cities == nil)
	fmt.Printf("%#v", cities)
	fmt.Print(len(cities))
	//cities[0] = "london" //error
	fmt.Println("")
	nums := []int{2, 3, 5, 6, 9, 6}
	fmt.Println(nums)
	fmt.Println("")
	var n = make([]int, 2)
	fmt.Printf("%#v\n", n)

	fmt.Println("")

	fmt.Println("")

	friends := []string{"DAN", "Maria"}
	fmt.Println(friends)

	fmt.Println("")

	type nm []string
	friendst := nm{"DAN", "Maria"}
	fmt.Println(friendst)
	friendst[0] = "RH"
	fmt.Println(friendst)

	for index, name := range friendst {
		fmt.Printf("%d: %q\n", index, name)
	}
	var nn []int
	fmt.Println("")
	fmt.Print(nn == nil)

	var m = []int{}
	fmt.Println("")
	fmt.Print(m == nil)

	a, b := []int{1}, []int{9}
	var eq bool = true
	for i, value := range a {
		if value != b[i] {
			eq = false
			break
		}

	}
	if len(a) != len(b) {
		eq = false
	}
	fmt.Println("")
	if eq == true {
		fmt.Print("EQULAL")
	} else {
		fmt.Print("NOT EQULAL")
	}
	numbers := []int{2, 5}
	numbers = append(numbers, 55)
	fmt.Println("")
	fmt.Printf("%#v", numbers)
	fmt.Println("")
	nmbr := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, nmbr...)
	fmt.Printf("%#v", numbers)

	fmt.Println("")
	var src = []int{1, 2, 3, 5}
	var dst = make([]int, len(src)-1)
	mmm := copy(dst, src)
	fmt.Println(src, dst, mmm)
	fmt.Println("")
}
