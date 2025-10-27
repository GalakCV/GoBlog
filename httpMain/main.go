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

	BlogHandlers := handlers.NewBlogHandler()
	BlogServerMux.HandleFunc("/", BlogHandlers.Home)

	
	// 4. Inicie o servidor com o SEU MUX
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