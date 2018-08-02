package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Kubernetes still rocks from master (v2)!\n")
                fmt.Println("User requested: /")
	})

	http.ListenAndServe(":8080", nil)
}
