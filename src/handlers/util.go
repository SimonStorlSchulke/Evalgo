package handlers

import (
	"fmt"
	"io/ioutil"
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
		return false
	}

	requester, err := user.FromMatrikel(loggedInMat)
	if err != nil {
		return false
	}

	if requester.Matrikel == us.Matrikel || courseconfig.GetConfig().Open_course || requester.IsAuthorized() {
		return true
	}

	if us.IsAuthorized() {
		return true
	}

	return false
}

func existingTaskNumbers() []int {

	list := make([]int, 0)
	tasks, err := ioutil.ReadDir("./Userdata/assignments")
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range tasks {
		fName := f.Name()
		fNumStr := fName[5 : len(fName)-3]
		fNum, _ := strconv.Atoi(fNumStr)
		list = append(list, fNum)
	}
	return list
}
