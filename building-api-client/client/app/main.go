package main

import (
	"log"

	"github.com/kings5layer/API-PROGRAMING/building-api-client/client"
)

func main() {
	c := client.NewClient("http://localhost:8888")
	if err := c.Login("user", "password"); err != nil {
		log.Fatal(err)
	}
}
