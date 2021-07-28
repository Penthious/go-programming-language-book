package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	year   int  `json:"released"` // not marshaled
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool hand Luke", year: 1962, Color: true, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
}

func main() {
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", data) //either use "%s" or fmt.Sprint() or strings(data) since data is a []bytes

	var titles []struct{ Title string }

	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	fmt.Println(titles)

}
