package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("App server starting...")
	var requestCount int
	requestCount = 0

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request from %s\n", r.RemoteAddr)

		requestCount++
		if requestCount > 5 {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Some internal error has occurred! This is Pod %s\n", hostname)
			return
		}

		fmt.Fprintf(w, "This is v3 running in pod %s\n", hostname)
	})

	http.ListenAndServe(":8080", nil)

}
