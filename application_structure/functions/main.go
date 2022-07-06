package main

import (
	"fmt"
	"strings"
)

func main() {
	nums := []float64{1, 2, 3, 4, 55.77, 66}
	var s, m = sumAndProduct(nums...)
	fmt.Println(s, m)
	names := []string{"Rana", "Hamid"}
	var ss = personalInformation(10, names...)
	fmt.Println(ss)
}

func sum(a, b int) (s int) {
	s = a + b
	return s
}

func f1(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}
func f2(a ...int) {
	a[0] = 50
}

func sumAndProduct(a ...float64) (s float64, m float64) {
	ss := 0.0
	mm := 1.0
	for _, v := range a {
		ss = ss + v
		mm = mm * v
	}
	return ss, mm
}
func personalInformation(age int, name ...string) string {
	names := strings.Join(name, " ")
	fmt.Println(names)
	var rt = fmt.Sprintf("age: %d, full name: %s", age, names)
	return rt
}
