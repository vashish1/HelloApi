package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

var port = os.Getenv("PORT")

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":"+port, nil)
}
