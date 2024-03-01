package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/page", page)
	http.HandleFunc("/file", file)

	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello page\n")
}

func page(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, header := range headers {
			fmt.Fprintf(w, "%v %v", name, header)
		}
	}
}

func file(w http.ResponseWriter, req *http.Request) {
	data, err := os.ReadFile("./files/test.txt")
	if err != nil {
		fmt.Fprint(w, "Error: can't read file to show the content")
	}
	fmt.Fprintf(w, string(data))
}
