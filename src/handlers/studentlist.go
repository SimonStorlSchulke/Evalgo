package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"../user"
)

//returns true and matrikel if password and matrikel match
func loggedIn(r *http.Request) (bool, int) {
	var session, cErr = r.Cookie("session")

	if cErr == nil {
		sessionSplitted := strings.Split(session.Value, "<split>")
		storedMat, err := strconv.Atoi(sessionSplitted[0])
		if err != nil {
			return false, 0
		}
		storedPw := sessionSplitted[1]
		currentUser, err := user.FromMatrikel(storedMat)

		if currentUser.GetPassword() == storedPw {
			return true, storedMat
		}
	}
	return false, 0
}

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
	loggedIn, mat := loggedIn(r)
	if loggedIn {
		currentUser, err = user.FromMatrikel(mat)
	}

	var portraits []string
	for _, st := range studentlist {
		portraits = append(portraits, st.GetPortraitPath())
	}

	pageData := struct {
		Nav         string
		Students    []user.Student
		CurrentUser user.Student
		Portraits   []string
	}{
		Nav:         getNav(),
		Students:    studentlist,
		CurrentUser: currentUser,
		Portraits:   portraits,
	}

	tmpl.Execute(w, pageData)
}
