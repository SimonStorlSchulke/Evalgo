package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"../courseconfig"
	"../user"
	"github.com/gorilla/mux"
)

func studentFromURL(r *http.Request) (user.User, error) {
	//Get Matrikel from URL
	params := mux.Vars(r)
	matrikelSt := params["matrikel"]
	matrikel, _ := strconv.Atoi(matrikelSt)

	student, err := user.FromMatrikel(matrikel)

	return student, err
}

func checkViewPermission(us user.User, r *http.Request) bool {

	loggedIn, loggedInMat := loggedIn(r)
	if !loggedIn {
		fmt.Println("1")
		return false
	}

	requester, err := user.FromMatrikel(loggedInMat)
	if err != nil {
		fmt.Println("2")
		return false
	}

	if requester.Matrikel == us.Matrikel || courseconfig.GetConfig().Open_course || requester.IsAuthorized() {
		return true
	}
	return false
}
