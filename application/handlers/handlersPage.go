package handlers

//Aqui vou somente realizar os handlers para os metodos GET para carregamento das Paginas

import (
	"net/http"
	"text/template"
)

type BlogStruct struct{}

func NewBlogHandler() *BlogStruct {
	return &BlogStruct{}
}

func (h *BlogStruct) Home(w http.ResponseWriter, r *http.Request){

    w.Header().Set("Content-Type", "text/html")
    w.Header().Set("X-Frame-Options", "DENY")
    
    if r.Method != "GET" {
        w.Header().Set("Allow", "GET")
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    files := []string{
        "views/templates/pages/base.html",
        "views/templates/pages/home.html",
    }
    template, err := template.ParseFiles(files...)
    if err != nil {
        http.Error(w,"Error parsing template", http.StatusInternalServerError)
        return
    }
    err = template.ExecuteTemplate(w, "base", nil)
    if err != nil {
        http.Error(w, "Error executing template", http.StatusInternalServerError)
    }
}

func (h *BlogStruct) Login(w http.ResponseWriter, r *http.Request){
    
    w.Header().Set("Content-Type", "text/html")
    w.Header().Set("X-Frame-Options", "DENY")
    
    if r.Method != "GET" {
        w.Header().Set("Allow", "GET")
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    files := []string{
        "views/templates/pages/base.html",
        "views/templates/pages/login.html",
    }
    template, err := template.ParseFiles(files...)
    if err != nil {
        http.Error(w,"Error parsing template", http.StatusInternalServerError)
        return
    }
    err = template.ExecuteTemplate(w, "base", nil)
    if err != nil {
        http.Error(w, "Error executing template", http.StatusInternalServerError)
    }

}

func (h *BlogStruct) SignUpGet(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "text/html")
    w.Header().Set("X-Frame-Options", "DENY")
    
    if r.Method != "GET" {
        w.Header().Set("Allow", "GET")
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    files := []string{
        "views/templates/pages/base.html",
        "views/templates/pages/signup.html",
    }
    template, err := template.ParseFiles(files...)
    if err != nil {
        http.Error(w,"Error parsing template", http.StatusInternalServerError)
        return
    }
    err = template.ExecuteTemplate(w, "base", nil)
    if err != nil {
        http.Error(w, "Error executing template", http.StatusInternalServerError)
    }
}