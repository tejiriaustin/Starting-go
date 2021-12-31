package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./static/index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Starting go server...")
}

func handleRequests() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", start)
	router.HandleFunc("./", indexHandler)
	router.HandleFunc("/get_user", get_User).Methods("GET")
	router.HandleFunc("/post_user/{name}/{id}", post_User).Methods("POST")
	router.HandleFunc("/delete_user/{name}/{id}", delete_User).Methods("DELETE")
	router.HandleFunc("/put_user/{name}/{id}", update_User).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8001", router))

}

func main() {
	createFile()
	handleRequests()
}
