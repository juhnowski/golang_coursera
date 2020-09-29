package main

import (
	"fmt"
	"os"
)

type name struct {
	fname string
	lname string
}

func main() {
	var fn string
	fmt.Println("Input name of the text file")
	fmt.Scan(&fn)

	file, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var f string
	var l string
	sli := make([]name, 0)

	for {
		_, err := fmt.Fscanf(file, "%s %s", &f, &l)
		if err != nil {
			break
		}
		n := name{fname: f, lname: l}
		sli = append(sli, n)
	}

	for _, v := range sli {
		fmt.Printf("%20s %20s \n", v.fname, v.lname)
	}
}
