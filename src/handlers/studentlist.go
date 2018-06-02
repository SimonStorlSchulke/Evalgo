package handlers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"../user"
)

func HandleStudents(w http.ResponseWriter, r *http.Request) {
	//Read jsons:
	studentlist := user.ReadStudents()

	//apply group colors
	c1, c2, c3, cDef := "#beffa3", "#a3e6ff", "#f7ffa8", "#808080"
	for i, _ := range studentlist {
		switch {
		case i%3 == 0:
			studentlist[i].Gruppenfarbe = c1
		case i%3 == 1:
			studentlist[i].Gruppenfarbe = c2
		case i%3 == 2:
			studentlist[i].Gruppenfarbe = c3
		default:
			studentlist[i].Gruppenfarbe = cDef
		}
	}

	tmpl, err := template.ParseFiles("./templates/studentlist.go.html")
	if err != nil {
		log.Fatalln(err)
	}

	var currentUser user.Student
	var session, cErr = r.Cookie("session")
	if cErr == nil {
		mat, _ := strconv.Atoi(session.Value)
		currentUser, err = user.FromMatrikel(mat)
	}

	pageData := struct {
		Nav         string
		Students    []user.Student
		CurrentUser user.Student
	}{
		Nav:         getNav(),
		Students:    studentlist,
		CurrentUser: currentUser,
	}

	tmpl.Execute(w, pageData)
}
