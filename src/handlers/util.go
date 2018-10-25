package handlers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"../courseconfig"
	"../user"
	"github.com/gorilla/mux"
)

//Messaging
const (
	MsgPermissionDeniedAuth     = "Permission Denied. This Page is only accessible by authorized (Tutor or Admin), logged in users."
	MsgPermissionDeniedLoggedIn = "Permission Denied. This Page is only accessible for logged in users."
	MsgPermissionDenied         = "Permission Denied."
	MsgErrorReadingMat          = "Error reading Matrikelnumber"
	MsgNoRessources             = "there are no Files in the coursedata/ressources directory"
	MsgMessageNotPost           = "Message is not Post"
)

//Prints a message to the users Browser
func WriteMsg(w http.ResponseWriter, str string) {
	//Todo: Format nicely
	fmt.Fprintf(w, str)
}

//The same as the function above but with an additional error
func WriteError(w http.ResponseWriter, str string, err error) {
	WriteMsg(w, fmt.Sprintf("An Error occured - %s: %v", str, err))
}

//Check Permission and return false if no loggged in user or user is not authorized
func isAuthSession(r *http.Request, w http.ResponseWriter) bool {
	loggedIn, loggedInMat := loggedIn(r)
	loggedInUser, err := user.FromMatrikel(loggedInMat)

	if !loggedIn || loggedInUser.Usertype == user.STUDENT || err != nil {
		return false
	}
	return true
}

//Returns student struct from current url
func studentFromURL(r *http.Request) (user.User, error) {
	//Get Matrikel from URL
	params := mux.Vars(r)
	matrikelSt := params["matrikel"]
	matrikel, _ := strconv.Atoi(matrikelSt)

	student, err := user.FromMatrikel(matrikel)

	return student, err
}

//Check whether user is permited to see a post or not.
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

//A list of ints that tasks have been created form in /coursedata/tasks directory
func existingTaskNumbers() []int {

	list := make([]int, 0)
	tasks, err := ioutil.ReadDir("./coursedata/tasks")
	if err != nil {
		fmt.Println(err)
		return list
	}

	for _, f := range tasks {
		if f.IsDir() {
			continue
		}
		fName := f.Name()
		fNumStr := fName[5 : len(fName)-3]
		fNum, _ := strconv.Atoi(fNumStr)
		list = append(list, fNum)
	}
	return list
}

//The highest number, that a Task has been created for in /coursedata/tasks directory
func highestTaskNumber() int {
	maxNr := 1
	for _, e := range existingTaskNumbers() {
		if e > maxNr {
			maxNr = e
		}
	}
	return maxNr
}

//Check if directory is empty
func isEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // not empty or error
}
