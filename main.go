package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
)


var todos = []string{
	"Fix the sump pump discharge pipes",
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo string
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
	log.Fatal(http.ListenAndServe(":8080", router))
}