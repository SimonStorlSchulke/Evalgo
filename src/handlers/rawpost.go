package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"../user"
	"github.com/gorilla/mux"
)

//Handlefunc for accessing raw unformated markdown posts
func HandleRawPosts(w http.ResponseWriter, r *http.Request) {

	//Get matrikel from URL
	params := mux.Vars(r)
	matrikelSt := params["matrikel"]
	matrikel, _ := strconv.Atoi(matrikelSt)

	//Display Error Message when matrikel does not exist
	student, err := user.FromMatrikel(matrikel)
	if err != nil {
		fmt.Fprintf(w, "Error reading Matrikelnumber %v for accessing raw post", matrikel)
		return
	}
	postNrSt := params["postnr"]
	postNr, _ := strconv.Atoi(postNrSt)

	//Parse Markdown to []byte
	content := student.GetPost(postNr)
	fmt.Fprint(w, string(content[:]))

}
