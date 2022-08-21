package main

import (
	"fmt"
	"math"
)

func main() {
	var X int
	var N int
	fmt.Scan(&X)
	fmt.Scan(&N)
	if N == 7 {
		fmt.Println(X)
	} else if N == 0 {
		fmt.Println(20)
	} else {
		XAlt := math.Max(0, float64(X-N))
		fmt.Println(XAlt)
	}
}
