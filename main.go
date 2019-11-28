package main

import (
	"fmt"
	"github.com/braulio94/base_project/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/project", controllers.GetProject).Methods("GET")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Serving on port", port)
}
