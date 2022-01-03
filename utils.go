package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type form struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var parent []form

func get_User(w http.ResponseWriter, r *http.Request) {
	byteArray, _ := readAsform()
	arr, _ := json.MarshalIndent(byteArray, "", " ")
	fmt.Fprintf(w, string(arr))
}

func delete_User(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Delete")
	vars := mux.Vars(r)
	id := vars["id"]

	byteArray, _ := readAsform()

	for index, i := range byteArray {
		if i.Id == id {
			byteArray = append(byteArray[:index], byteArray[index+1:]...)
		}
	}
	store, _ := json.Marshal(byteArray)
	openAndWrite(store)
}

func update_User(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Update")
	delete_User(w, r)
	parent, _ = readAsform()
	fmt.Println(parent)
	var new_form form

	new_form.Name = r.FormValue("name")
	new_form.Id = r.FormValue("id")

	byteArray, _ := readAsform()
	byteArray = append(byteArray, new_form)

	arr, _ := json.MarshalIndent(byteArray, "", " ")
	openAndWrite(arr)
}

func post_User(w http.ResponseWriter, r *http.Request) {
	parent, _ = readAsform()
	fmt.Println(parent)
	var new_form form

	new_form.Name = r.FormValue("name")
	new_form.Id = r.FormValue("id")

	byteArray := append(parent, new_form)

	arr, _ := json.MarshalIndent(byteArray, "", " ")

	fmt.Fprintf(w, string(arr))
	openAndWrite(arr)
	fmt.Println("has been written!")
}
