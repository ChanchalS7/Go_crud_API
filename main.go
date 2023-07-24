package main
import (
	"fmt"
	"github.com/gorilla/mux"
	"todo/routes"
	"log"
	"net/http"
)

func main(){
	router := mux.NewRouter()
	routes.SetRoutes(router)
	fmt.Println("API Server is running...")	
	log.Fatal(http.ListenAndServe(":8080", router))
}