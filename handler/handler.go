package handler

import (
	"fmt"
	"net/http"

	"example.com/go-api/api"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the homepage!"))
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("About page"))
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusMethodNotAllowed)
		return
	}
	question := r.Form.Get("question")
	fmt.Fprintf(w, "Your question is , %s", question)
}

func AskQuestion(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusMethodNotAllowed)
		return
	}
	question := r.Form.Get("question")
	fmt.Fprintf(w, "Your question is  %s", question)
	//fmt.Println("Your question is ", question)
	api.SendQuestion(question)
}
