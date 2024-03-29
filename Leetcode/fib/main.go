package main

import "fmt"

func main() {
	fmt.Print(fib(13))
}
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	//var result = [10]int{}
	var result = make([]int, n+1)
	result[0] = 0
	result[1] = 1
	for i := 2; i <= n; i++ {
		result[i] = result[i-2] + result[i-1]
	}
	return result[n]
}
