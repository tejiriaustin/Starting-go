package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./static/index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func handleRequests() {

	router := mux.NewRouter().StrictSlash(true)

	port := os.Getenv("PORT")
	fmt.Println(port)

	fmt.Print("Starting go server...")
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/post_user", post_User).Methods("POST")
	router.HandleFunc("/get_user", get_User).Methods("GET")
	router.HandleFunc("/delete_user/{name}/{id}", delete_User).Methods("DELETE")
	router.HandleFunc("/put_user/{name}/{id}", update_User).Methods("PUT")

	log.Fatal(http.ListenAndServe(":"+port, router))

}

func main() {
	createFile()
	handleRequests()
}
