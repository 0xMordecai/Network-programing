package main

import (
	"fmt"
	"net/http"
)

type ShoppingList struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Items []string `json:"items"`
}

var allData []ShoppingList

func main() {
	http.HandleFunc("POST /v1/lists", handleCreateList)
	fmt.Println("listening on port :8888")
	http.ListenAndServe(":8888", nil)
}
