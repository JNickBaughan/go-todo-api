package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
)

// Todo is the basic type to hold a todo item
type todo struct {
	ID string `json:"ID"`
    ParentID string `json:"ParentID"`
	Desc string `json:"Desc"`
	Complete bool `json:"Complete"`
}

var todos []todo

func addTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo todo
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Uh Oh... something went wrong there")
	}
	
	json.Unmarshal(reqBody, &newTodo)
	todos = append(todos, newTodo)
	w.WriteHeader(http.StatusCreated)
	
	json.NewEncoder(w).Encode(newTodo)
}

func listTodos(w http.ResponseWriter, r *http.Request) {
	todoList, err := json.Marshal(todos)
	if err != nil {
		fmt.Fprintf(w, "Uh Oh... looks like something went wrong")
	}
	fmt.Fprintf(w, string(todoList))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", listTodos)

	router.HandleFunc("/add", addTodo)
	log.Fatal(http.ListenAndServe(":3000", router))
}