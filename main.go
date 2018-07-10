package main

import (
	"fmt"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Kubernetes still rocks\n")
                fmt.Println("User requested: /")
	})

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
