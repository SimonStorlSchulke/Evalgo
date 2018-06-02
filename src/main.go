package main

import (
	"net/http"

	"./handlers"
	"github.com/gorilla/mux"
)

func main() {

	rtr := mux.NewRouter()

	FileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", FileServer))

	FileServerPortraits := http.FileServer(http.Dir("Userdata/Portraits"))
	http.Handle("/portraits/", http.StripPrefix("/portraits/", FileServerPortraits))

	rtr.HandleFunc("/register", handlers.HandleRegister)
	rtr.HandleFunc("/login", handlers.HandleLogin)
	rtr.HandleFunc("/", handlers.HandleStudents)
	rtr.HandleFunc("/post", handlers.HandlePost)
	rtr.HandleFunc("/profile/{matrikel}", handlers.HandleProfile)
	http.Handle("/", rtr)
	http.ListenAndServe(":8080", nil)

}
