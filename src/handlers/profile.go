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

func HandleProfile(w http.ResponseWriter, r *http.Request) {

	//Get Matrikel from mux Parameters
	params := mux.Vars(r)
	matrikelSt := params["matrikel"]
	matrikel, _ := strconv.Atoi(matrikelSt)
	student, err := user.FromMatrikel(matrikel)
	if err != nil {
		fmt.Println("Error reading Matrikelnumber")
	}

	//Parse Markdown to []byte
	md := blackfriday.MarkdownCommon(student.GetPosts())

	pageData := struct {
		St      user.Student
		Nav     string
		Profile string
	}{
		St:      student,
		Nav:     getNav(),
		Profile: string(md[:]),
	}

	tpl := template.Must(template.ParseFiles("./templates/profile.go.html"))
	tpl.Execute(w, pageData)
}
