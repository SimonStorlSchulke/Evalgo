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
	matrikel, _ := strconv.ParseInt(r.FormValue("matrikel")[0:], 10, 64)

	page := map[string]string{
		"nav":  getNav(),
		"name": vorname,
	}

	tpl.Execute(w, page)
	user.NewStudent(r.FormValue("vorname"), nachname, int(matrikel)).Register()
}
