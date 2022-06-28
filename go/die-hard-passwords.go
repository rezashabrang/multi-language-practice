package main

import (
	"fmt"
	"math"
	"strconv"
)

var PRIMES []int

func main() {
	var n int
	fmt.Scan(&n)
	PRIMES = append(PRIMES, 2)
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
	if number < 2 {
		return false
	}
	for _, prime := range PRIMES {
		if number%prime == 0 {
			return true
		}
	}
	for i := PRIMES[len(PRIMES)-1]; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
		if checkPrime(i) {
			PRIMES = append(PRIMES, i)
		}
	}
	return true

}
