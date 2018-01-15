package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%v%v", "Hello World from ", hostname)
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
