package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// fmt.Println("Path", os.Args[0])
	// fmt.Println("Arg", os.Args[1])
	// fmt.Println("2nd Arg", os.Args[2])
	// fmt.Println("N: ", len(os.Args))

	var result, err = strconv.Atoi("5.3")
	if err == nil {
		fmt.Print(result)
	} else {
		fmt.Println(err)
	}
	var s = map[string]int{
		"a": 1,
	}
	_ = s
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	var j = 10
	for j >= 10 {
		fmt.Printf("Hello %d", j)
		j--
	}
	fmt.Println()
	var count = 0
	for i := 0; true; i++ {
		if i%13 == 0 {
			count++
			fmt.Printf("%d is divisible ", i)
		}
		if count == 10 {
			break
		}
	}
	fmt.Println("JUST")
	fmt.Println()
	people := [5]string{"Helen", "Mark", "Brenda", "Antonio", "Micheal"}
	friends := [2]string{"Mark", "Brenda"}

	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("index %d, firend %q", index, friend)
				break
			}
		}
	}
	fmt.Println()
	for j, name := range people {
		fmt.Printf("%d=%s", j, name)
	}
	var i = 0
loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}
	for k, val := range friends {
		fmt.Printf("%d %q", k, val)
	}
	fmt.Println()
	language := "Python"
	switch language {
	case "python":
		fmt.Println("you are learning python")
	case "go", "golang":
		fmt.Println("Go is popular")
	default:
		fmt.Println("any language is good")
	}

	var n = 10
	switch true {
	case n%2 == 0:
		fmt.Println("any language is good")
	case n%2 == 1:
		fmt.Println("Go is popular")
	default:
		fmt.Println("any language is good-96")
	}

	hour := time.Now().Hour()
	fmt.Println(hour)
	switch {
	case hour < 12:
		fmt.Println("good morning")
	case hour < 17:
		fmt.Println("Good evening")
	default:
		fmt.Println("Good evening")
	}
}
