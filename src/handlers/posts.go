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

	us, err := studentFromURL(r)
	if err != nil {
		fmt.Println("Error reading Matrikelnumber")
	}

	//Check Permission
	if !checkViewPermission(us, r) {
		fmt.Fprintf(w, "Permission Denied")
		return
	}

	params := mux.Vars(r)
	postNrSt := params["postnr"]
	postNr, _ := strconv.Atoi(postNrSt)

	//Parse Markdown to []byte
	md := blackfriday.MarkdownCommon(us.GetPost(postNr))

	//bool to check if there is Feedback in Template
	FbIs := false

	fb, err := user.GetFeedback(us.Matrikel, postNr)
	if err == nil {
		FbIs = true
	} else {
		fmt.Println(err)
	}

	pageData := struct {
		St      user.User
		Profile string
		PostNr  int
		FbIs    bool
		Fb      user.Feedback
	}{
		St:      us,
		Profile: string(md[:]),
		PostNr:  postNr,
		FbIs:    FbIs,
		Fb:      fb,
	}

	tpl := template.Must(template.ParseFiles("./templates/posts.go.html"))
	tpl.Execute(w, pageData)
}
