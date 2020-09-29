package main

import (
	"fmt"
	"strconv"
)

func main() {
	var snum string

	fmt.Println("Input float number")
	fmt.Scan(&snum)

	i1, err := strconv.ParseFloat(snum, 64)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println(int(i1))
	}

}
