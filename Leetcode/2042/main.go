package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	var v = areNumbersAscending("sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s")
	fmt.Println((v))

}
func areNumbersAscending(s string) bool {
	var words = strings.Split(s, " ")
	var minVal = math.MinInt64
	for i := 0; i < len(words); i++ {
		//fmt.Printf("%d=%s, ", i, words[i])
		if val, err := strconv.Atoi(words[i]); err == nil {

			if val <= minVal {
				return false
			} else {
				minVal = val
			}
		}
	}
	return true
}
