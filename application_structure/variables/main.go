package main

import "fmt"

func main() {
	var age int = 30
	fmt.Println("Age: ", age)
	var name = "Rana Hamid"
	_ = name
	//fmt.Println("Your name is", name)
	s := "Learnign GO Lang!"

	fmt.Println(s)
	var t = 5
	fmt.Println(t)
	car, cost := "Audi", 5000
	fmt.Println(car, cost)
	car, newCost := "honda", 120
	_ = newCost

	var (
		salary    float64
		gender    bool
		firstName string
	)
	fmt.Println(salary, gender, firstName)
	var i, j int
	i, j = 5, 8
	fmt.Println(i, j)

	var a = 4
	var b = 5.2
	fmt.Println(a, b, firstName)

	//a = int(b)
	//fmt.Println(a, b, firstName)
	/*
		i = 0
		j = 0
		fmt.Println(i, j)
	*/
	fmt.Printf("\"Hello GO!\"\n%.2f\n", b)
	var figure = "Circle"
	var radius = 5
	var pi = 3.141592654
	fmt.Printf("Radisu is =%d. area =%s =%f", radius, figure, 2*pi*float64(radius))
	fmt.Print("Radius\n")
	fmt.Printf("This is a %v", figure)
	fmt.Printf("This is a %T", figure)
	fmt.Printf("This is a %T", pi)
	fmt.Printf("This is a %T", radius)
	closed := true
	fmt.Printf("File is closed %T", closed)
	fmt.Printf("55 is in binary %08b\n", 55)
	fmt.Printf("%.3f", 5.563554)
}
