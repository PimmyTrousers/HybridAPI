package main

import (
	"fmt"
	"log"
)

func main() {
	args, err := ParseArgs()
	if err != nil {
		log.Fatal(err)
	}
	config, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", config)
	fmt.Printf("%v\n", args)
}
