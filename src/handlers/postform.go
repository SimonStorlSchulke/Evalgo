package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"../user"
)

//Handles user Post form
func HandlePostForm(w http.ResponseWriter, r *http.Request) {
	//Redirect to loggin if not logged in
	isLoggedIn, matrikel := loggedIn(r)
	if isLoggedIn == false {
		http.Redirect(w, r, "./login", http.StatusSeeOther)
	}
	tpl := template.Must(template.ParseFiles("./templates/postform.go.html"))

	postedText := r.FormValue("postedText")
	tm := time.Now()
	timestring := fmt.Sprintf("<i class='last-modified'>last modified %v:%v %s %v %v</i>\n", tm.Hour(), tm.Minute(), tm.Month(), tm.Day(), tm.Year())
	postedText = timestring + postedText
	postNumber, _ := strconv.Atoi(r.FormValue("postNr"))
	st, _ := user.FromMatrikel(matrikel)
	if postedText != "" && postNumber > 0 {
		st.PostNr(postedText, postNumber)
		redirPath := fmt.Sprintf("./%v/post/%v", matrikel, postNumber)
		http.Redirect(w, r, redirPath, http.StatusSeeOther)
	}

	tpl.Execute(w, st)
}
