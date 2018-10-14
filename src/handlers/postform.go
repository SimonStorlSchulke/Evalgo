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
	//get highest taskNumber
	if isLoggedIn == false {
		http.Redirect(w, r, "./login", http.StatusSeeOther)
	}
	tpl := template.Must(template.ParseFiles("./templates/postform.go.html"))

	us, _ := user.FromMatrikel(matrikel)
	pageData := struct {
		Us         user.User
		MaxPostNum int
	}{
		Us:         us,
		MaxPostNum: highestTaskNumber(),
	}

	tpl.Execute(w, pageData)
}

//Handle incoming Posts
func ProcessPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprint(w, "Method is not Post")
		return
	}
	fmt.Println("2")
	isLoggedIn, matrikel := loggedIn(r)
	if !isLoggedIn {
		http.Redirect(w, r, "./login", http.StatusSeeOther)
		return
	}

	us, _ := user.FromMatrikel(matrikel)
	postcontent := r.FormValue("postcontent")
	postNr, err := strconv.Atoi(r.FormValue("postNr"))

	if err != nil {
		fmt.Fprint(w, "cannot convert PostNumber to int", err)
		return
	}
	err = us.PostNr(postcontent, postNr)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	redirURL := fmt.Sprintf("./?nr=%v&mat=%v", postNr, matrikel)
	http.Redirect(w, r, redirURL, http.StatusSeeOther)
}
