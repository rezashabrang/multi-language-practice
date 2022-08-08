package main

import (
	"fmt"
)

func main() {
	var number int
	var numberList = []int{}
	fmt.Scan(&number)
	for number != 0 {
		numberList = append(numberList, number)
		fmt.Scan(&number)
	}
	for i := 0; i < len(numberList); i++ {
		fmt.Println(numberList[len(numberList) - i - 1])
	}
}
