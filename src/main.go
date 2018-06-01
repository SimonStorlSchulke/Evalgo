package main

import (
	"fmt"
	"net/http"

	"./handlers"
	"./user"
)

func main() {

	//serve static files (css, js...)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println(fs)

	//handle profiles TODO -> nicht ganze Liste lesen, nur Matrikelnummern
	stList := user.ReadStudents()
	for _, st := range stList {
		http.HandleFunc(fmt.Sprintf("/profile/%v", st.Matrikel), handlers.HandleProfile)
	}

	http.HandleFunc("/register", handlers.HandleRegister)
	http.HandleFunc("/studentlist", handlers.HandleStudents)
	http.HandleFunc("/profile", handlers.HandleProfile)
	http.ListenAndServe(":8080", nil)
}
