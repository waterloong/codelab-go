package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.URL.Query()["user"]
	if ok {
		fmt.Fprintf(w, "Hello, %s!", user[0])
	} else {
		fmt.Fprintf(w, "Hello, world!")
	}
}
