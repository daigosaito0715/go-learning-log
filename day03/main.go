package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello API")
}

func main() {

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server Start : http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server Error:", err)
	}
}
