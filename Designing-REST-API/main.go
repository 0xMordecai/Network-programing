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
type ShoppingListPatch struct {
	Name  *string  `json:"name"`
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
	//	The update endpoint
	http.HandleFunc("PUT /v1/lists/{id}", handleUpdateList)
	//	The Patch endpoint
	http.HandleFunc("PATCH /v1/lists/{id}", handlePatchList)
	//	The retriver endpoint
	http.HandleFunc("GET /v1/lists/{id}", handleGetList)
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

func handleUpdateList(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	for i, list := range allData {
		if strconv.Itoa(list.ID) == id {
			var updatedList ShoppingList
			err := json.NewDecoder(r.Body).Decode(&updatedList)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			allData[i] = updatedList
			if err := json.NewEncoder(w).Encode(updatedList); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	}
	http.Error(w, "List not found", http.StatusNotFound)
}

func handlePatchList(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	for i, list := range allData {
		if strconv.Itoa(list.ID) == id {
			var patch ShoppingListPatch
			err := json.NewDecoder(r.Body).Decode(&patch)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if patch.Name != nil {
				list.Name = *patch.Name
			}
			if patch.Items != nil {
				list.Items = patch.Items
			}
			allData[i] = list
			err = json.NewEncoder(w).Encode(list)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	}
	http.Error(w, "List not found", http.StatusNotFound)
}

func handleGetList(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	for _, list := range allData {
		data, err := json.Marshal(list)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
