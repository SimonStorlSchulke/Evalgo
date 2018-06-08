package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"../user"
)

//Use md5 to Hash password
func HashPassword(pw string) string {
	hasher := md5.New()
	hasher.Write([]byte(pw))
	return hex.EncodeToString(hasher.Sum(nil))
}

//compare unhashed and hashed password and return true if they match
func CheckPasswordHash(enteredPw, hashedPw string) bool {
	if HashPassword(enteredPw) == hashedPw {
		return true
	} else {
		return false
	}
}

//Returns true and matrikel if password and matrikel match
func loggedIn(r *http.Request) (bool, int) {
	var session, cErr = r.Cookie("session")

	if cErr == nil {
		sessionSplitted := strings.Split(session.Value, "<split>")
		storedMat, err := strconv.Atoi(sessionSplitted[0])
		if err != nil {
			return false, 0
		}
		storedPw := sessionSplitted[1]
		currentUser, err := user.FromMatrikel(storedMat)

		if currentUser.GetPassword() == storedPw {
			return true, storedMat
		}
	}
	return false, 0
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {

	tpl := template.Must(template.ParseFiles("./templates/login.go.html"))
	//create Student from Matrikelnumber and Redirect if not existing
	matrikel, _ := strconv.ParseInt(r.FormValue("matrikel")[0:], 10, 64)
	var err error
	var us user.User

	//dirty Fix. Clean me up pls
	if matrikel != 0 {
		us, err = user.FromMatrikel(int(matrikel))
		if err != nil {
			fmt.Println(err)
		}
		hashedPassword := us.GetPassword()
		enteredPassword := r.FormValue("password")

		//set session if password correct
		if CheckPasswordHash(enteredPassword, hashedPassword) {
			cookieValue := fmt.Sprintf("%v<split>%s", strconv.Itoa(int(matrikel)), hashedPassword)
			c := &http.Cookie{
				Name:  "session",
				Value: cookieValue,
			}
			http.SetCookie(w, c)
			http.Redirect(w, r, "./", http.StatusSeeOther)
		} else {
			fmt.Print("user entered wrong password")
		}
	}
	tpl.Execute(w, "")
}
