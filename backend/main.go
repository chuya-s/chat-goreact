package main

import (
	"fmt"
	"net/http"
)

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Server start")
	})
}

func main() {
	setupRoutes()
	http.ListenAndServe("localhost:8000", nil)
}
