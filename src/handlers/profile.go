package handlers

import (
	"net/http"
	"path"
	"strconv"
	"text/template"

	"../user"
	"github.com/russross/blackfriday"
)

func HandleProfile(w http.ResponseWriter, r *http.Request) {

	//Get Matrikel from Url
	matrikel, _ := strconv.Atoi(path.Base(r.URL.Path))
	testSt := user.FromMatrikel(matrikel)

	//Parse Markdown to []byte
	md := blackfriday.MarkdownCommon(testSt.GetPosts())

	pageData := struct {
		St      user.Student
		Nav     string
		Profile string
	}{
		St:      testSt,
		Nav:     getNav(),
		Profile: string(md[:]),
	}

	tpl := template.Must(template.ParseFiles("./templates/profile.go.html"))
	tpl.Execute(w, pageData)
}
