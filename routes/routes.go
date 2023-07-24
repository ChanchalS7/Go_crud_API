package routes

import (
	"github.com/gorilla/mux"
	"todo/controllers"
	"net/http"
)

//set Routes sets up the API routes
func SetRoutes(router *mux.Router){
	router.HandleFunc("/api/todos", controllers.GetTodos).Methods("GET")
	router.HandleFunc("/api/todos/{id}", controllers.GetTodoByID).Methods("GET")
	router.HandleFunc("/api/todos",controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/api/todos/{id}",controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("api/todos/{id}",controllers.DeleteTodo).Methods("DELETE")
	http.Handle("/",router)
}