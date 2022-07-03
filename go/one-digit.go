package main

import (
	"fmt"
	"strconv"
)

func main() {
	var stringDigits string
	fmt.Scan(&stringDigits)
	for {
		var digits = []rune(stringDigits)
		var digitsSum int = 0
		for i := 0; i < len(stringDigits); i++ {
			digit, _ := strconv.Atoi(string(digits[i]))
			digitsSum += digit

		}
		if digitsSum < 10 {
			fmt.Println(digitsSum)
			break
		}
		stringDigits = strconv.Itoa(digitsSum)
	}
}
