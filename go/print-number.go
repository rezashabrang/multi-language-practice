package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number string
	fmt.Scan(&number)

	// Appending digits to an array
	var digits = []rune(number)
	for i := 0; i < len(number); i++ {
		digit, _ := strconv.Atoi(string(digits[i]))
		fmt.Printf("%v: ", digit)
		for j := 0; j < digit; j++ {
			fmt.Print(digit)
		}
		fmt.Println()
	}

}
