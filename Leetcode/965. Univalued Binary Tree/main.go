package main

import "math"

func main() {

}
func findComplement(num int) int {
	var sum = 0

	for counter := 0; true; counter++ {
		var bit = num & 1
		if bit == 0 {
			sum = sum + int(math.Pow(2, float64(counter)))
		}
		num = num >> 1
		if num <= 0 {
			break
		}

	}
	return sum
}
