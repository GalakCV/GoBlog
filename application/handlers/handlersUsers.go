package handlers

//Aqui, sao as rotas de API

import (
	"fmt"
	"net/http"
)

type UserStruct struct{}

func NewUsersHandler() *UserStruct {
	return &UserStruct{}
}

func UserSignUp(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil{
		http.Error(w, "Error processing input", http.StatusBadRequest)
		return
	}
	Email := r.FormValue("email")
	Password := r.FormValue("password")

	fmt.Fprintf(w, "Email: %s, Senha: %s", Email, Password)
}
