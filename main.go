package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/filters", FilterIndex)
	router.HandleFunc("/filters/{filterId}", FilterShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func FilterIndex(w http.ResponseWriter, r *http.Request) {
	filters := Filters{
		Filter{Name: "filter name"},
		Filter{Name: "another filter"},
	}

	json.NewEncoder(w).Encode(filters)
}

func FilterShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filterId := vars["filterId"]
	fmt.Fprintln(w, "Filter show:", filterId)
}

type Filter struct {
	Name         string
	CreatedTime  time.Time
	ModifiedTime time.Time
}

type Filters []Filter
