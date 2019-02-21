package main

import (
	"net/http"
	"os"
        "fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "[{\"size\":\"n/a\",\"description\":\"not available\"}]")
	})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}


