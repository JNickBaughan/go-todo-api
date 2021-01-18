package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
	"go-todo-api/models"
)

func indexOf(id string, todos []models.Todo) int {
	for i, t := range todos {
		if  id == t.ID {
			return i
		}
	}
	return -1    //not found.
 }

 func removeTodo(i int, todos []models.Todo) []models.Todo {
	copy(todos[i:], todos[i+1:]) // Shift a[i+1:] left one index.
	return todos[:len(todos)-1] 
 }
 
var todos []models.Todo

func addTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo models.Todo
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

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["id"]
	
	todos = removeTodo(indexOf(todoID, todos), todos)

	updatedList, e := json.Marshal(todos)
	
	if e != nil {
		fmt.Fprintf(w, "Uh Oh... something went wrong when trying to delete that todo")
	}

	fmt.Fprintf(w, string(updatedList))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", listTodos)

	router.HandleFunc("/add", addTodo)

	router.HandleFunc("/delete/{id}", deleteTodo)


	log.Fatal(http.ListenAndServe(":3000", router))
}