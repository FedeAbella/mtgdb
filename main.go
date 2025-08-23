package main

import (
	"fmt"
	"log"
)

func main() {

	data, err := ReadSetCards()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", data.Data[:5])
}
