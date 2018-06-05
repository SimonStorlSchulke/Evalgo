package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"../user"
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
)

//Handlefunc to display all Posts of a student
func HandleProfile(w http.ResponseWriter, r *http.Request) {

	//Get Matrikel from mux Parameters
	params := mux.Vars(r)
	matrikelSt := params["matrikel"]
	matrikel, _ := strconv.Atoi(matrikelSt)
	student, err := user.FromMatrikel(matrikel)
	if err != nil {
		fmt.Fprintf(w, "There is no Student with the Matrikel %v registered.", matrikel)
		return
	}

	//Parse Markdown to []byte
	postData, postNumbers := student.GetAllPosts()
	md := blackfriday.MarkdownCommon(postData)

	pageData := struct {
		St          user.Student
		Nav         string
		Profile     string
		PostNumbers []int
	}{
		St:          student,
		Nav:         getNav(),
		Profile:     string(md[:]),
		PostNumbers: postNumbers,
	}

	tpl := template.Must(template.ParseFiles("./templates/profile.go.html"))
	tpl.Execute(w, pageData)
}
