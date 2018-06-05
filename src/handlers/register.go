package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strconv"
	"text/template"

	"../user"
)

func HashPassword(pw string) string {
	hasher := md5.New()
	hasher.Write([]byte(pw))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CheckPasswordHash(enteredPw, hashedPw string) bool {
	if HashPassword(enteredPw) == hashedPw {
		return true
	} else {
		return false
	}
}

//Handles Register Form and Submit
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/register.go.html"))

	vorname := r.FormValue("vorname")
	nachname := r.FormValue("nachname")
	passwort := HashPassword(r.FormValue("passwort"))
	matrikel, _ := strconv.ParseInt(r.FormValue("matrikel")[0:], 10, 64)

	err := user.NewStudent(r.FormValue("vorname"), nachname, int(matrikel), passwort).Register()

	//Redirect if Registration successfull
	if err == nil {
		http.Redirect(w, r, "./", http.StatusSeeOther)
	}
	tpl.Execute(w, vorname)
}
