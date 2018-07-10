package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Kubernetes rocks v2!")
                fmt.Println("Successfully responded to internet citizen")
	})

	http.ListenAndServe(":8080", nil)
}
