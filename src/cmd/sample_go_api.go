package main

import (
	"log"
	"net/http"

	"github.com/marythought/sample_go_api/src/filter"
)

func main() {

	router := filter.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
