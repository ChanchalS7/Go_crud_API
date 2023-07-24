package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"todo/models"
)

var todos []models.Todo
var currentID int

//GET Todos return all todo items
func GetTodos( w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(todos)
}

//GetTodoByID returns a todo item by its ID

func GetTodoByID( w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id,_:= strconv.Atoi(params["id"])
	for _,todo:= range todos {
		if todo.ID==id{
			json.NewEncoder(w).Encode(todo)
			return 
		}
	}
	json.NewEncoder(w).Encode(nil)
}

// CreateTodo creates a new todo item
func CreateTodo( w http.ResponseWriter, r*http.Request){
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	currentID++;
	todo.ID= currentID
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}


// UpdateTodo updates an existing todo item

func UpdateTodo(w http.ResponseWriter, r *http.Request){

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var updatedTodo models.Todo 
	_ = json.NewDecoder(r.Body).Decode(&updatedTodo)
	for i, todo := range todos{
		if todo.ID == id {
			todos[i] = updatedTodo
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}

//DeleteTodo deletes an existing todo by its ID

func DeleteTodo(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, todo := range todos{
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}











