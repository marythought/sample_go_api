package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!\n")
}

func FilterIndex(w http.ResponseWriter, r *http.Request) {
	filters := Filters{
		Filter{Name: "filter name"},
		Filter{Name: "another filter"},
	}

	if err := json.NewEncoder(w).Encode(filters); err != nil {
		panic(err)
	}
}

func FilterShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filterId := vars["filterId"]
	fmt.Fprintln(w, "Filter show:", filterId)
}
