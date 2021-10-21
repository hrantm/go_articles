package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/articles/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Return Signle Article Endpoint")
	vars := mux.Vars(r)
	key := vars["id"]
	var result Article
	for _, article := range Articles {
		if article.Id == key {
			result = article
			break
		}
	}
	json.NewEncoder(w).Encode(result)

}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all articles endpoint")
	json.NewEncoder(w).Encode(Articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage")
	fmt.Println("Hit Homepage")
}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello2", Desc: "desc 2", Content: "Content 2"},
	}
	fmt.Println("Starting Server")
	handleRequests()
}
