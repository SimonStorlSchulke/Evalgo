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
	rtr.HandleFunc("/post", handlers.HandlePostForm)
	rtr.HandleFunc("/info", handlers.HandleInfo)
	rtr.HandleFunc("/{number}", handlers.HandleStudents)
	rtr.HandleFunc("/", handlers.HandleStudents)
	rtr.HandleFunc("/profile/{matrikel}", handlers.HandleProfile)
	rtr.HandleFunc("/{matrikel}/postraw/{postnr}", handlers.HandleRawPosts)
	rtr.HandleFunc("/{matrikel}/post/{postnr}", handlers.HandlePosts)
	http.Handle("/", rtr)
	http.ListenAndServe(":8080", nil)
}
