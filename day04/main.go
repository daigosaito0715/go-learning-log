package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	user := User{
		Name: "daigo",
		Job:  "backend",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {

	http.HandleFunc("/user", userHandler)

	fmt.Println("Server Start : http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}