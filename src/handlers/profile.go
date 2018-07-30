package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"../user"
	"github.com/russross/blackfriday"
)

//Handlefunc to display all Posts of a student
func HandleProfile(w http.ResponseWriter, r *http.Request) {

	//Get Matrikel from mux Parameters
	student, err := studentFromURL(r)

	//Check session
	if !checkViewPermission(student, r) {
		fmt.Fprintf(w, "Permission Denied")
		return
	}

	if err != nil {
		fmt.Fprintf(w, "There is no Student with the Matrikel %v registered.", student.Matrikel)
		return
	}

	//Parse Markdown to []byte
	postData, postNumbers := student.GetAllPosts()
	md := blackfriday.MarkdownCommon(postData)

	pageData := struct {
		St          user.User
		Profile     string
		PostNumbers []int
	}{
		St:          student,
		Profile:     string(md[:]),
		PostNumbers: postNumbers,
	}

	tpl := template.Must(template.ParseFiles("./templates/profile.go.html"))
	tpl.Execute(w, pageData)
}
