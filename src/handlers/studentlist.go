package handlers

import (
	"log"
	"net/http"
	"text/template"

	"../user"
)

func HandleStudents(w http.ResponseWriter, r *http.Request) {
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

	/*
		//TODO
		page := struct {
			nav      string
			students []user.Student
		}{
			"test",
			studentlist,
		}
	*/

	tmpl, err := template.ParseFiles("./templates/studentlist.go.html")
	if err != nil {
		log.Fatalln(err)
	}
	tmpl.Execute(w, studentlist)
}
