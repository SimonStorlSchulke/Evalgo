package handlers

import (
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

	//Redirect if Registration successfull
	if err == nil {
		http.Redirect(w, r, "./", http.StatusSeeOther)
		//TODO: Login directly
	}
	tpl.Execute(w, vorname)
}
