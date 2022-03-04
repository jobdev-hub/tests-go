package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {

	// get all albums
	res, err := http.Get("http://localhost:8080/albums")
	if err != nil {
		log.Fatal(err.Error())
	}

	// decode json response (jsonToStruct)
	var albums []album
	err = json.NewDecoder(res.Body).Decode(&albums)
	if err != nil {
		log.Fatal(err.Error())
	}

	// print all albums by struct
	fmt.Println(albums)

	// editable print all albums by struct
	for _, album := range albums {
		log.Printf("%s by %s: $%.2f\n", album.Title, album.Artist, album.Price)
	}

	// encode struct to json
	toJson, err := json.Marshal(albums)
	if err != nil {
		log.Fatal(err.Error())
	}

	// print json
	fmt.Println(string(toJson))
}
