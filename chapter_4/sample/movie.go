package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	type Movie struct {
		Title  string
		Year   int  `json:"released"`
		Color  bool `json:"color,omitempty"`
		Actors []string
	}

	var movies = []Movie{

		{Title: "aaaa", Year: 1942, Color: false, Actors: []string{"dddd", "bbbb"}},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON Marsharling faild: %s", data)
	}

	fmt.Printf("%s\n", data)

	data, err = json.MarshalIndent(movies, "", "    ")

	if err != nil {
		log.Fatalf("JSON Marsharling faild: %s", data)
	}
	fmt.Printf("%s\n", data)

}
