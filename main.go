package main

//Added goFMT file watcher
//added .gitignore for .idea
//removed additional libraries

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Good Uncle")
	})

	mux.HandleFunc("/hello/golang", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Food for college!")
	})

	http.ListenAndServe(":9000", mux)
}
