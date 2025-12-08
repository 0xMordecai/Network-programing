package main

import (
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
}
