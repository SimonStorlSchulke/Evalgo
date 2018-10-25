package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"../user"
)

//Handles user Post form
func HandlePostForm(w http.ResponseWriter, r *http.Request) {

	//Redirect to loggin if not logged in
	isLoggedIn, matrikel := loggedIn(r)
	if isLoggedIn == false {
		http.Redirect(w, r, "./login", http.StatusSeeOther)
		return
	}
	tpl := template.Must(template.ParseFiles("./templates/postform.go.html"))

	us, err := user.FromMatrikel(matrikel)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	pageData := struct {
		Us         user.User
		MaxPostNum int
	}{
		Us:         us,
		MaxPostNum: highestTaskNumber(),
	}
	err = tpl.Execute(w, pageData)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	return
}

//Handle incoming Posts
func ProcessPost(w http.ResponseWriter, r *http.Request) {

	//Max Memory 10kb
	r.ParseMultipartForm(10 * 1024)

	if r.Method != "POST" {
		WriteMsg(w, MsgMessageNotPost)
		return
	}

	isLoggedIn, matrikel := loggedIn(r)
	if !isLoggedIn {
		http.Redirect(w, r, "./login", http.StatusSeeOther)
		return
	}
	us, err := user.FromMatrikel(matrikel)
	if err != nil {

		WriteError(w, "", err)
		return
	}
	postcontent := r.FormValue("postcontent")
	postNr, err := strconv.Atoi(r.FormValue("postNr"))

	if err != nil {
		WriteError(w, "cannot convert PostNumber to int", err)
		return
	}
	err = us.PostNr(postcontent, postNr)
	if err != nil {
		WriteError(w, "", err)
		return
	}
	redirURL := fmt.Sprintf("./?nr=%v&mat=%v", postNr, matrikel)
	http.Redirect(w, r, redirURL, http.StatusSeeOther)
}
