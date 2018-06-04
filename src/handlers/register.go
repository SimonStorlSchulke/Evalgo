package handlers

import (
	"net/http"
	"strconv"
	"text/template"

	"../user"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//Handles Register Form and Submit
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/register.go.html"))

	vorname := r.FormValue("vorname")
	nachname := r.FormValue("nachname")
	passwort, _ := HashPassword(r.FormValue("passwort"))
	matrikel, _ := strconv.ParseInt(r.FormValue("matrikel")[0:], 10, 64)

	err := user.NewStudent(r.FormValue("vorname"), nachname, int(matrikel), passwort).Register()

	//Redirect if Registration successfull
	if err == nil {
		http.Redirect(w, r, "./", http.StatusSeeOther)
	}
	tpl.Execute(w, vorname)
}
