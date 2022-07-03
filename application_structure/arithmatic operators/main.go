package main

import (
	"fmt"
	"strconv"
)

type km float64
type mile float64

func main() {

	a, b := 4, 2
	r := (a + b) / (a - b) * 2
	fmt.Println(r)
	r = 5 % 2
	r++
	fmt.Println(r)

	x := 1
	x++
	x += 1
	fmt.Println(x > a)
	fmt.Println(false && true)

	var z uint8 = 255
	fmt.Println(z)
	z++
	fmt.Println(z)

	s := string(99)
	fmt.Println(s)

	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)
	var myStr1 = fmt.Sprintf("%d", 34244)
	fmt.Println(myStr1)

	fmt.Println(string(34234))
	s1 := "3.1235"
	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Println(f1)
	fmt.Printf("%T\n", f1)

	i, err := strconv.Atoi("-50")
	fmt.Println(i)
	k := strconv.Itoa(-50)
	fmt.Println(k)
	//underlying type

	type speed uint
	var ss1 speed = 10
	var ss2 speed = 20
	fmt.Println(ss2 - ss1)

	var xx int = 10
	var s3 speed = speed(xx)
	fmt.Println(s3)
	/// km
	var paristolondon km = 465
	var distsc km = 452
	ditnace := paristolondon / distsc
	fmt.Println(ditnace)
	var d rune = 5
	fmt.Println(d)

	type second = uint
	var hour second = 3600
	fmt.Printf("%d", hour/60)
}
