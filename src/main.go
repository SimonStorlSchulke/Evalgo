package main

import (
	"html/template"
	"net/http"
	"strconv"

	"./user"
)

func handleStudents(w http.ResponseWriter, r *http.Request) {
	//Read jsons:
	studentlist := user.ReadStudents()

	//apply group colors
	red, green, blue, grey := "#ed4b4b", "#66ed4b", "#55b2f4", "#808080"
	for i, _ := range studentlist {
		switch {
		case i%3 == 0:
			studentlist[i].Gruppenfarbe = red
		case i%3 == 1:
			studentlist[i].Gruppenfarbe = green
		case i%3 == 2:
			studentlist[i].Gruppenfarbe = blue
		default:
			studentlist[i].Gruppenfarbe = grey
		}
	}

	tmpl, _ := template.ParseFiles("studentlist.go.html")
	tmpl.Execute(w, studentlist)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("register.go.html"))

	vorname := r.FormValue("vorname")
	nachname := r.FormValue("nachname")
	matrikel, _ := strconv.ParseInt(r.FormValue("matrikel")[0:], 10, 64)

	tpl.Execute(w, vorname)
	user.NewStudent(r.FormValue("vorname"), nachname, int(matrikel)).Register()
}

func main() {

	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/studentlist", handleStudents)
	http.ListenAndServe(":1313", nil)
}
