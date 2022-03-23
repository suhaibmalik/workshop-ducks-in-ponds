package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// load file
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}

	// debug
	fmt.Println(string(data))

	// todo: Marshal
	// v := ??
	// err = json.Unmarshal(data, &v)

	// todo: compare ponds and ducks

	// todo: print results
}
