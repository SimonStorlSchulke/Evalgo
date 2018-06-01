package main

import (
	"net/http"

	"./handlers"
)

func HandleProfiles() {

}

func main() {

	http.HandleFunc("/register", handlers.HandleRegister)
	http.HandleFunc("/studentlist", handlers.HandleStudents)
	http.HandleFunc("/profile", handlers.HandleProfile)
	http.ListenAndServe(":1313", nil)
}
