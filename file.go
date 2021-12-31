package main

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
	"os"
)

func createFile() {
	database, err := os.OpenFile("DataBase.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening the database: %v", err)
	}
	defer database.Close()
}
func openAndWrite(value []byte) {

	error := ioutil.WriteFile("DataBase.txt", value, 0666)
	if error != nil {
		log.Print(error)
	}
}

func openAndRead() ([]byte, error) {
	b, err := ioutil.ReadFile("Database.txt")
	if err != nil {
		log.Print(err)
	}
	return b, err
}

func readAsform() ([]form, error) {
	var u []form
	arr, _ := openAndRead()
	err := binary.Read(bytes.NewBuffer(arr[:]), binary.BigEndian, &u)
	return u, err
}
