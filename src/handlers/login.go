package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"../user"
)

//Returns true and matrikel if password and matrikel match
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

func HandleLogin(w http.ResponseWriter, r *http.Request) {

	tpl := template.Must(template.ParseFiles("./templates/login.go.html"))

	//create Student from Matrikelnumber and Redirect if not existing
	matrikel, _ := strconv.ParseInt(r.FormValue("matrikel")[0:], 10, 64)
	var st user.Student
	var err error

	//dirty Fix. Clean me up pls
	if matrikel != 0 {
		st, err = user.FromMatrikel(int(matrikel))
		if err != nil {
			fmt.Println(err)
		}
		password := st.GetPassword()
		enteredPassword := r.FormValue("password")

		//set session if password correct
		if password == enteredPassword {
			cookieValue := fmt.Sprintf("%v<split>%s", strconv.Itoa(int(matrikel)), password)
			c := &http.Cookie{
				Name:  "session",
				Value: cookieValue,
			}
			http.SetCookie(w, c)
			http.Redirect(w, r, "./", http.StatusSeeOther)
		} else {
			fmt.Print("user entered wrong password")
		}
	}
	tpl.Execute(w, "")
}
