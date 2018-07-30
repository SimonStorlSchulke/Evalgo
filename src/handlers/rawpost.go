package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Handlefunc for accessing raw unformated markdown posts
func HandleRawPosts(w http.ResponseWriter, r *http.Request) {

	//Get matrikel from URL
	us, err := studentFromURL(r)

	//Check session
	if !checkViewPermission(us, r) {
		fmt.Fprintf(w, "Permission Denied")
		return
	}

	//Display Error Message when matrikel does not exist
	if err != nil {
		fmt.Fprintf(w, "Error reading Matrikelnumber %v for accessing raw post", us.Matrikel)
		return
	}
	postNrSt := mux.Vars(r)["postnr"]
	postNr, _ := strconv.Atoi(postNrSt)

	//Parse Markdown to []byte
	content := us.GetPost(postNr)
	fmt.Fprint(w, string(content[:]))

}
