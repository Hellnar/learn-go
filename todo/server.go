package main

import (
	"fmt"
	"net/http"
)

var state []string

func main() {
	http.HandleFunc("/", getTasks)
	http.HandleFunc("/add", addTask)

	http.ListenAndServe(":8090", nil)
}

func getTasks(w http.ResponseWriter, req *http.Request) {
	length := len(state)
	fmt.Fprintf(w, "<p>You have %v tasks", length)
	return
}

func addTask(w http.ResponseWriter, req *http.Request) {
	task := req.URL.Query().Get("task")
	fmt.Println(task)
	if task != "" {
		state = append(state, task)
	}
	fmt.Fprint(w, "Successfully added a task")
}
