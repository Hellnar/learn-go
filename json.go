package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type game struct {
	Name    string
	Players int
	Seasons []string
}

func main() {
	data, err := os.ReadFile("./files/test.json")
	if err != nil {
		panic(err)
	}

	var game map[string]interface{}
	if err := json.Unmarshal(data, &game); err != nil {
		panic(err)
	}

	fmt.Println(game)
	comp := game["name"]
	fmt.Println(comp)
	seasons := game["seasons"]
	fmt.Println(seasons)
}
