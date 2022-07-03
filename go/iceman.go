package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n>100 {
		fmt.Println("Steam")
	} else if n<0 {
		fmt.Println("Ice")
	} else {
		fmt.Println("Water")
	}

}
