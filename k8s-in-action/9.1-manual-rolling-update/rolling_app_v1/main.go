package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("App server starting...")
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request from %s\n", r.RemoteAddr)

		fmt.Fprintf(w, "This is v1 running in pod %s\n", hostname)
	})

	http.ListenAndServe(":8080", nil)

}
