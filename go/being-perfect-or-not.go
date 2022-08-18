package main

import (
	"fmt"
	"math"
)

func main() {
	var number int
	var dividerSum int = 0
	fmt.Scan(&number)

	checkPoint := math.Ceil(float64(number) / 2)
	for i := 2; i <= int(checkPoint); i++ {
		if number%i == 0 {
			dividerSum += i
		}

	}
	if dividerSum + 1 == number {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
