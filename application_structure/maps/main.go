package main

import (
	"fmt"
)

func main() {
	var m = map[string]int{
		"A": 65,
	}
	fmt.Println(m)
	fmt.Println(len(m))
	var n = map[string]int{}
	fmt.Printf("%#v\n", n)
	fmt.Println(len(n))

	fmt.Println(m["A"])
	fmt.Println(m["B"])

	m["B"] = 66
	fmt.Println(m["B"])

	map1 := make(map[int]int)
	map1[1] = 10
	m["B"] = 660
	fmt.Println(m)
	fmt.Println(len(m))
	m["C"] = 100
	value, ok := m["C"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("No value for m[c]")
	}
	for k, v := range m {
		fmt.Printf("%s=%d\n", k, v)
	}

	for k, v := range m {
		fmt.Printf("%#v=%#v\n", k, v)
	}
	delete(m, "C")
	for k, v := range m {
		fmt.Printf("%s=%d\n", k, v)
	}

	mm := map[string]string{
		"A": "X",
	}

	nn := map[string]string{
		"A": "X",
	}
	s1 := fmt.Sprintf("%s", mm)
	s2 := fmt.Sprintf("%s", nn)

	fmt.Println(s1, s2)
	if s1 == s2 {
		fmt.Println("Both are same")
	} else {
		fmt.Println("Not same ")
	}

	friends := map[string]int{"Dan": 40, "RH": 32}
	var neighbours = friends
	neighbours["RH"] = 31
	fmt.Println(friends)
	fmt.Println(neighbours)
	people := map[string]int{} //make(map[string]int)
	for k, v := range friends {
		people[k] = v
	}
	friends["RH"] = 50
	friends["Hamid"] = 500
	fmt.Println(friends)
	fmt.Println(people)
}
