package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type form struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func get_User(w http.ResponseWriter, r *http.Request) {
	byteArray, err := openAndRead()
	if err != nil {
		log.Fatal("Error reading from file")
	}
	error := json.Unmarshal(byteArray, w)
	if error != nil {
		log.Print(err)
	}
	fmt.Printf("%s", w)
}

func delete_User(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Delete")
	vars := mux.Vars(r)
	id := vars["Id"]

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
	elem, _ := ioutil.ReadAll(r.Body)

	byteArray, _ := openAndRead()
	byteArray = append(byteArray, elem...)

	openAndWrite(byteArray)
}

func post_User(w http.ResponseWriter, r *http.Request) {
	byteArray, err := openAndRead()
	req_Body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading response")
	}
	byteArray = append(byteArray, req_Body...)

	openAndWrite(byteArray)
	fmt.Fprint(w, byteArray)
	fmt.Print("has been writtten to Database file")
}
