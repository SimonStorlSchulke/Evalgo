package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"../user"
)

//TODO
func alreadyLoggedIn(r *http.Request) bool {
	return false
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {

	//Check if already logged in
	/*_, err := r.Cookie("session")
	if err == nil {
		fmt.Println("ups1")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}*/

	tpl := template.Must(template.ParseFiles("./templates/login.go.html"))

	//create Student from Matrikelnumber and Redirect if not existing
	matrikel, _ := strconv.ParseInt(r.FormValue("matrikel")[0:], 10, 64)
	var st user.Student
	var err error

	//dirty Fix. Clean me up pls
	if matrikel != 0 {
		st, err = user.FromMatrikel(int(matrikel))
		if err != nil {
			fmt.Println(err)
		}
		password := st.GetPassword()
		enteredPassword := r.FormValue("password")

		//set session if password correct. TODO: check always when on site
		if password == enteredPassword {
			c := &http.Cookie{
				Name:  "session",
				Value: strconv.Itoa(int(matrikel)),
			}
			http.SetCookie(w, c)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			fmt.Print("user entered wrong password")
		}
	}

	page := map[string]string{
		"nav": getNav(),
	}

	tpl.Execute(w, page)
}
