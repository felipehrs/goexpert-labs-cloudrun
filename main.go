package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/felipehrs/go-expert-cloud-run/controller"
)

func main() {
	http.HandleFunc("/getinfo", func(w http.ResponseWriter, r *http.Request) {
		GetWeatherHandler(w, r, http.Get, http.Get)
	})
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
