package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	var name string
	var address string

	fmt.Println("Input name")
	fmt.Scan(&name)
	fmt.Println("Input address")
	fmt.Scan(&address)

	p := map[string]string{
		"name":    name,
		"address": address,
	}
	var jsonData []byte
	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))

}
