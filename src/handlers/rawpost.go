package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"../user"
	"github.com/gorilla/mux"
)

func HandleRawPosts(w http.ResponseWriter, r *http.Request) {

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
	content := student.GetPost(postNr)
	fmt.Fprint(w, string(content[:]))
}
