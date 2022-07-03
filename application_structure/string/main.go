package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println
	fmt.Println("He Says\"Hello\" ")

	fmt.Println(`He Says:\n "Hello" `)
	fmt.Println(`Price: 100
hi hello
what?`)
	fmt.Println(`c:\useres`)
	fmt.Println(`c:\useres`)
	fmt.Println()
	str := "Rana"
	for i, ch := range str {
		fmt.Printf("index: %d: %d, ", i, ch)
	}
	fmt.Println()
	for _, ch := range str {
		fmt.Printf("index: %c, ", ch)
	}
	fmt.Println("")
	fmt.Println(len("╛"))
	fmt.Println(len("R"))
	s1 := "রানা হামিদ"
	fmt.Println(s1[2:7])

	rs := []rune(s1)
	fmt.Printf("%T\n", rs)
	fmt.Println(string(rs[0:4]))

	var result = strings.Contains("I রানা love Go Programming!", "রানা")
	p(result)

	result = strings.ContainsAny("I রানা love Go Programming!", "gpj")
	p(result)

	result = strings.ContainsRune("I রানা love Go Programming!", 'g')
	p(result)

	var n = strings.Count("I রানা love Go Programming!", "g")
	p(n)

	var str8 = strings.ToLower("I রানা love Go Programming!")
	p(str8)

	str8 = strings.ToUpper("I রানা love Go Programming!")
	p(str8)
	p(strings.EqualFold("GP", "gp"))

	p(strings.Replace("I রানা love Go Programming!", "love", "hate", -2))

	p(strings.ReplaceAll("I রানা love Go Programming!", "love", "hate"))

	s := strings.Split("who am i? what you are", " ")
	fmt.Printf("%T\n", s)
	fmt.Printf("%#v\n", s)

	ss := strings.Join(s, "-")
	fmt.Printf("%T\n", ss)
	fmt.Printf("%#v\n", ss)

	s = strings.Split("who am i?\n what you are", "")

	fmt.Printf("%T\n", s)
	fmt.Printf("%#v\n", s)

	var fields = strings.Fields("who am i?\n what you are")
	fmt.Printf("%T\n", fields)
	fmt.Printf("%#v\n", fields)

	s1 = strings.TrimSpace("\t goodbyte window. welcome linux  ")
	fmt.Printf("%q\n", s1)

	s1 = strings.Trim("...Hello Golphers?!!!!How are you????", "?.!")
	fmt.Printf("%q\n", s1)
}
