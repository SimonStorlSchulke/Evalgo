package handlers

import (
	"net/http"
	"text/template"

	"../user"
)

//Handles user Post form
func HandlePost(w http.ResponseWriter, r *http.Request) {
	//Redirect to loggin if not logged in
	isLoggedIn, matrikel := loggedIn(r)
	if isLoggedIn == false {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	tpl := template.Must(template.ParseFiles("./templates/post.go.html"))

	postedText := r.FormValue("postedText")
	st, _ := user.FromMatrikel(matrikel)
	if postedText != "" {
		st.Post(postedText)
	}

	tpl.Execute(w, st)
}
