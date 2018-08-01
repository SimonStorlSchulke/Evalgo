package handlers

import (
	"net/http"
	"strconv"
	"text/template"

	"../user"
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
)

//Handle posts for /{matrikel}/post/{postnr}
func HandleTasks(w http.ResponseWriter, r *http.Request) {

	postNrSt := mux.Vars(r)["tasknr"]
	postNr, _ := strconv.Atoi(postNrSt)

	//Parse Markdown to []byte
	md := blackfriday.MarkdownCommon(user.GetTask(postNr))

	pageData := struct {
		Task   string
		PostNr int
	}{
		Task:   string(md[:]),
		PostNr: postNr,
	}

	tpl := template.Must(template.ParseFiles("./templates/tasks.go.html"))
	tpl.Execute(w, pageData)
}
