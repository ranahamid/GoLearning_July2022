package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

type address struct {
	name    string
	street  string
	city    string
	state   string
	PinCode int
}

func main() {

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Bob1", age: 21})
	fmt.Println(newPerson("Jon"))
	var s = person{name: "SS", age: 66}
	fmt.Println(s)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(s.age)
	fmt.Println(sp.age)

	var a address
	a.name = "Shah makdum avenue"
	a.street = "sect 13"
	fmt.Println(a)

	var b = &address{name: "01", street: "sec", city: "77", state: "55", PinCode: 444}
	fmt.Println((*b))

}
