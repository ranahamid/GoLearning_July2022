package main

import "fmt"

func main() {
	title1, author1, year1 := "The Devince Comedy", "Dante Aligheri", 1320
	title2, author2, year2 := "Macbeth", "William Shakespeare", 1606

	fmt.Println("Book1: ", title1, author1, year1)
	fmt.Println("Book1: ", title2, author2, year2)

	type book struct {
		title, author string
		year          int
	}

	type book1 struct {
		title, author string
		year          int
	}

	mybook := book{title: "The Devince Comedy", author: "Dante Aligheri", year: 1320}
	fmt.Println("Book1: ", mybook)
	_ = mybook
	aBook := book{title: "HII"}
	fmt.Printf("Book1: %#v\n", aBook)

	lastBoo := book{title: "Anna Karenina"}
	fmt.Printf("Book1: %#v\n", lastBoo)
	lastBoo.title = "EE"
	fmt.Printf("Book1: %#v\n", lastBoo)
	lastBoo.year = 1990
	fmt.Printf("Book1: %+v\n", lastBoo)
	fmt.Println(lastBoo == mybook)

	mybook2 := aBook
	aBook.author = "NEW YEAR "
	fmt.Printf("Book1: %+v\n", mybook2)

	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "diana",
		lastName:  "Muller",
		age:       30,
	}
	fmt.Printf("%#v\n", diana)
	fmt.Printf("%+v\n", diana)
	fmt.Printf("%d\n", diana.age)

	type Book struct {
		string
		float64
		bool
	}
	b1 := Book{"1985", 65.22, false}
	fmt.Printf("%#v", b1)

	type Employee struct {
		name  string
		salay int
		bool
	}
	e := Employee{"RH", 100, false}
	fmt.Printf("%#v\n", e)
	e.bool = true
	fmt.Printf("%#v\n", e)

	type Contact struct {
		email, address string
		phone          int
	}
	type Emp struct {
		name        string
		salay       int
		ContactInfo Contact
	}
	john := Emp{
		name:  "RH ",
		salay: 400,
		ContactInfo: Contact{
			email:   "rana@gmail.com",
			address: "hoouse 55, s77",
			phone:   0171447522,
		},
	}
	fmt.Printf("%#v\n", john)
	john.ContactInfo.email = "ranahamid007@gmail.com"
	fmt.Printf("%#v\n", john)
}
