package handlers

import (
	"net/http"
	"text/template"
)

func HandleProfile(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/profile.go.html"))
	test := "blablub heytataratata"
	tpl.Execute(w, test)
}
