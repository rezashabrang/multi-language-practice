package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	startSequence := math.Pow10(n - 1)
	endSequence := math.Pow10(n)
	for i := startSequence; i < endSequence; i++ {
		if checkDieHardness(int(i), n) {
			fmt.Println(i)
		}

	}

}

func checkDieHardness(number int, n int) bool {
	for i := 0; i < n; i++ {
		if !checkPrime(number) {
			return false
		}
		strNum := strconv.Itoa(number)
		number, _ = strconv.Atoi(strNum[:len(strNum)-1])

	}
	return true

}

func checkPrime(number int) bool {
	for i := 2; i <= int(math.Ceil(float64(number)/2)); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true

}
