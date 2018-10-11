package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"../user"
)

//Handles Register Form and Submit
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/register.go.html"))

	vorname := r.FormValue("vorname")
	nachname := r.FormValue("nachname")
	passwort := HashPassword(r.FormValue("passwort"))
	matrikel, _ := strconv.ParseInt(r.FormValue("matrikel")[0:], 10, 64)

	registeredUser := user.NewUser(r.FormValue("vorname"), nachname, int(matrikel), passwort)

	err := registeredUser.Register()

	//Redirect and autologin if Registration successfull
	if err == nil {

		//Auto Login afterwards
		cookieValue := fmt.Sprintf("%v<split>%s", strconv.Itoa(int(matrikel)), passwort)
		c := &http.Cookie{
			Name:  "session",
			Value: cookieValue,
		}
		http.SetCookie(w, c)
		http.Redirect(w, r, "./", http.StatusSeeOther)
	}
	tpl.Execute(w, vorname)
}
