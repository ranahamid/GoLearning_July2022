package main

import (
	"fmt"
	"unsafe"
)

func main() {

	s1 := []int{10, 20, 30, 40, 50}
	fmt.Println(s1)
	s3, s4 := s1[0:2], s1[1:3]
	fmt.Println(s3)
	fmt.Println(s4)

	carrs := []string{"HONDA", "Toyota"}
	fmt.Println(carrs)
	var newCars = []string{}
	newCars = append(newCars, carrs[0:2]...)

	var dest = make([]string, len(carrs))
	var n = copy(dest, carrs)

	carrs[0] = "changed"
	fmt.Println(newCars)
	fmt.Println(n)
	fmt.Println(dest)

	fmt.Println(len(s3), cap(s3))
	s3 = s1[1:2]
	fmt.Println(len(s3), cap(s3))
	fmt.Printf("%p\n", &s1)
	fmt.Printf("%p\n", &s3)

	fmt.Printf("%d\n", unsafe.Sizeof(s1))
	fmt.Printf("%d\n", unsafe.Sizeof(s3))
}
