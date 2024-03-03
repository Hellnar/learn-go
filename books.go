package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", home)

	http.ListenAndServe(":8090", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	data, err := os.ReadFile("./files/books.json")
	if err != nil {
		panic(err)
	}

	var books map[string]interface{}
	if err := json.Unmarshal(data, &books); err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(data))
}
