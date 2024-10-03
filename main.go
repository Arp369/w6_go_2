package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Message struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// GET Handler
func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid GET Method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Welcome to the GO API..")
}

// POST Handler
func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Post Method", http.StatusMethodNotAllowed)
		return
	}

	var msg Message
	//Read body of the post request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid read request Body", http.StatusBadRequest)
		return
	}

	//parse the JSON data
	err = json.Unmarshal(body, &msg)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	//response with the message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func main() {
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)
	fmt.Println("Server is running on post: 8990")
	log.Fatal(http.ListenAndServe(":8990", nil))
}
