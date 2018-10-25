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

	//Check Permission
	if !checkViewPermission(us, r) {
		fmt.Fprintf(w, MsgPermissionDenied)
		return
	}

	//Display Error Message when matrikel does not exist
	if err != nil {
		//Do not use WriteMsg() here - this handler only displays the raw post information and should therefore not contain any styling
		fmt.Fprintf(w, "%s %v for accessing raw post", MsgErrorReadingMat, us.Matrikel)
		return
	}
	postNrSt := mux.Vars(r)["postnr"]
	postNr, _ := strconv.Atoi(postNrSt)

	//Parse Markdown to []byte
	content := us.GetPost(postNr)
	fmt.Fprint(w, string(content[:]))
}
