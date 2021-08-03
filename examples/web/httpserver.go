// Write an http server which returns 'hello world' for any request.
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})
	http.ListenAndServe(":8080", nil)
}