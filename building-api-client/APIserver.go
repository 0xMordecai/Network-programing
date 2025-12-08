package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

var token string
var randomGen *rand.Rand

func random(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != "Bearer "+token {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{
		"value": randomGen.Intn(100),
	})
}

func setSeed(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != "Bearer "+token {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var data map[string]int
	json.NewDecoder(r.Body).Decode(&data)
	randomGen = rand.New(rand.NewSource(int64(data["seed"])))
}
