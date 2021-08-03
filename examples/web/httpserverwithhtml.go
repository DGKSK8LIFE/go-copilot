// Write an http web server that returns an html header that shows the user's IP address and the time.
// Usage:
// go run examples/web/httpserverwithhtml.go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<html><body><h1>Your IP is %s. The time is %s</h1></body></html>",
		r.RemoteAddr, time.Now().Format(time.RFC1123))
}