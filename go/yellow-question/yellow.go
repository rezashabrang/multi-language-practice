package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	os := strings.Repeat("o", n)
	fmt.Printf("W%vw!", os)

}
