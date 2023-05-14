package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type GlobalResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type CategoriesResponse struct {
	Status string     `json:"status"`
	Data   []Category `json:"data"`
}

type Category struct {
	Id          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"Description"`
}

func homePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, _ := json.Marshal(GlobalResponse{Status: "ok", Message: "This is homepage"})
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	fmt.Println("hit the homepage")
}

func about(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, _ := json.Marshal(GlobalResponse{Status: "ok", Message: "This is about"})
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	fmt.Println("hit the about")
}

func getAllCategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Categories := []Category{
		{Id: 1, Title: "Science", Description: "Science books"},
		{Id: 2, Title: "Fiction", Description: "Fiction books"},
	}
	res, _ := json.Marshal(CategoriesResponse{Status: "ok", Data: Categories})
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	fmt.Println("404 not found")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(GlobalResponse{Status: "nok", Message: "404 not found"})
	w.WriteHeader(http.StatusNotFound)
	w.Write(res)
	fmt.Println("404 not found")
}

func handleRequests() {
	router := httprouter.New()
	router.GET("/", homePage)
	router.GET("/about", about)
	router.GET("/categories", getAllCategories)
	router.NotFound = http.HandlerFunc(notFound)
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	handleRequests()
}
