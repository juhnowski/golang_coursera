package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 3)

	var s string
	for s != "X" {
		fmt.Println("enter integer")
		fmt.Scan(&s)
		if s == "X" {
			break
		}

		i, err := strconv.Atoi(s)
		if err == nil {
			sli = append(sli, i)
			sort.Ints(sli)
			fmt.Println(sli)
		}
	}

}
