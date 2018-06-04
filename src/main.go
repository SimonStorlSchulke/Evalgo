package main

import (
	"fmt"
	"net/http"

	"./handlers"
	"github.com/gorilla/mux"
)

func main() {

	conf := handlers.GetConfig()
	//TODO: rootUrl := conf.Root_url
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

	fmt.Printf("Start server at port%s:\n"+
		"   Coursname: %s\n"+
		"   Number of studentgroups: %v\n"+
		"   Course open: %v:",
		conf.Port, conf.Course_name, conf.Group_number, conf.Open_course)

	http.ListenAndServe(conf.Port, nil)
}
