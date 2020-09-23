package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	responseFunction := func(w http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write(body)

	}

	http.HandleFunc("/", responseFunction)
	http.ListenAndServe(":8080", nil)

}
