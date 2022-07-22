package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i < n+1; i++ {
		for j := 1; j < n + 1; j++ {
			fmt.Printf("%v ", i * j)
		}
		fmt.Print("\n")
	}

}
