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

//Handle posts for /{matrikel}/post/{postnr}
func HandlePosts(w http.ResponseWriter, r *http.Request) {

	//Get Matrikel from URL
	params := mux.Vars(r)
	matrikelSt := params["matrikel"]
	matrikel, _ := strconv.Atoi(matrikelSt)

	student, err := user.FromMatrikel(matrikel)
	if err != nil {
		fmt.Println("Error reading Matrikelnumber")
	}

	postNrSt := params["postnr"]
	postNr, _ := strconv.Atoi(postNrSt)

	//Parse Markdown to []byte
	md := blackfriday.MarkdownCommon(student.GetPost(postNr))

	pageData := struct {
		St      user.Student
		Nav     string
		Profile string
		PostNr  int
	}{
		St:      student,
		Nav:     getNav(),
		Profile: string(md[:]),
		PostNr:  postNr,
	}

	tpl := template.Must(template.ParseFiles("./templates/posts.go.html"))
	tpl.Execute(w, pageData)
}
