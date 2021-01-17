package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "(version: %s):Hello from go webserver on path %q", os.Getenv("VERSION"), html.EscapeString(r.URL.Path) )
}

func main() {
    http.HandleFunc("/", greetHandler)
    log.Fatal(http.ListenAndServe("0.0.0.0:7777", nil))
}
