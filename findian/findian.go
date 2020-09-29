package main

import (
	"fmt"
	"strings"
)

func main() {

	var s string

	fmt.Println("Input string")
	fmt.Scan(&s)

	ls := strings.ToLower(s)

	if strings.HasPrefix(ls, "i") && strings.HasSuffix(ls, "n") && strings.Index(ls, "a") != -1 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
