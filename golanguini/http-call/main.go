package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://api.openbrewerydb.org/v1/breweries"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

	breweries := parseBreweriesResponse(bytes)
	fmt.Println(breweries)
}

func parseBreweriesResponse(content []byte) []Brewery {
	breweries := make([]Brewery, 0, 20)
	err := json.Unmarshal(content, &breweries)
	if err != nil {
		panic(err)
	}

	return breweries
}

type Brewery struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude,string"`
	Latitude  float64 `json:"latitude,string"`
}
