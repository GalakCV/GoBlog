package main

import (
	"fmt"
	"goblog/application/handlers"
	"net/http"
)

func StartServer() {

	BlogServerMux := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	BlogServerMux.Handle("/static/", http.StripPrefix("/static/", fs))

	//rotas - GET
	BlogHandlers := handlers.NewBlogHandler()
	BlogServerMux.HandleFunc("GET /home", BlogHandlers.Home)
	BlogServerMux.HandleFunc("GET /login", BlogHandlers.Login)
	BlogServerMux.HandleFunc("GET /signup", BlogHandlers.SignUpGet)

	//rotas - POST
	BlogServerMux.HandleFunc("POST /api/register", handlers.UserSignUp)

	
	//server mux iniciando..
	fmt.Println("Starting server on port 8000...")
	fmt.Println("Server running at http://localhost:8000")
	err := http.ListenAndServe(":8000", BlogServerMux)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func main() {
	StartServer()
}