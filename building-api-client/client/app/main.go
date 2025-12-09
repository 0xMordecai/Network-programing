package main

import (
	"fmt"
	"log"

	"github.com/kings5layer/API-PROGRAMING/building-api-client/client"
)

func main() {
	c := client.NewClient("http://localhost:8888")
	if err := c.Login("user", "password"); err != nil {
		log.Fatal(err)
	}

	value, err := c.Random()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Random value:", value)

	if err = c.SetSeed(42); err != nil {
		log.Fatal(err)
	}

	value, err = c.Random()
	if err != nil {
		log.Fatal(err)
	}
}
