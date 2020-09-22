package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	responseFunction := func(w http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)

		if err != nil {
			log.Fatal(err)
		}

		w.Write(body)

	}

	http.HandleFunc("/", responseFunction)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
