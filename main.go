package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type RequestBody struct {
	Message string
}

func main() {
	h1 := func(w http.ResponseWriter, req *http.Request) {
		var b RequestBody
		err := json.NewDecoder(req.Body).Decode(&b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		io.WriteString(w, b.Message)
	}

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
