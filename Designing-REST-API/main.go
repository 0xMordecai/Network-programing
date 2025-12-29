package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ShoppingList struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Items []string `json:"items"`
}

var allData []ShoppingList

func main() {
	// The creation endpoint
	http.HandleFunc("POST /v1/lists", handleCreateList)
	// The list endpoint
	http.HandleFunc("GET /v1/lists", handleListLists)
	// The delete endpoint
	http.HandleFunc("DELETE /v1/lists/{id}", handleDeleteList)
	fmt.Println("listening on port :8888")
	http.ListenAndServe(":8888", nil)
}

func handleCreateList(w http.ResponseWriter, r *http.Request) {
	var list ShoppingList
	//	Unmarshall request's body
	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//	if everything went well ,we will store information in our allData var and return the newly created instance
	allData = append(allData, list)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleListLists(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(allData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleDeleteList(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	for i, list := range allData {
		if strconv.Itoa(list.ID) == id {
			allData = append(allData[:i], allData[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "List not found", http.StatusNotFound)
}
