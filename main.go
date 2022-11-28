package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.PathPrefix("/public").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/add-my-project", addMyproject).Methods("GET")
	route.HandleFunc("/form-add-project", FormAddProject).Methods("POST")
	route.HandleFunc("/detail-project", DetailProject).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")

	fmt.Println("Server running on localhost:5100")
	http.ListenAndServe("localhost:5100", route)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/home.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, nil)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func addMyproject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/add-project.html")

	if tmpl == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Message : " + err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, nil)
	}
}
func DetailProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/detail-project.html")

	if tmpl == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Message : " + err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, nil)
	}
}
func FormAddProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
    
	fmt.Println("Project Name : " + r.PostForm.Get("project-name")) // tag name
	fmt.Println("Star Date : " + r.PostForm.Get("start-date"))
	fmt.Println("End Date : " + r.PostForm.Get("end-date"))
	fmt.Println("Description : " + r.PostForm.Get("description"))
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}