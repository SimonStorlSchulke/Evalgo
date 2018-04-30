package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"

	"./user"
)

func handleStudents(w http.ResponseWriter, r *http.Request) {
	//Read json:
	jsondata, err := ioutil.ReadFile("./students.json")
	if err != nil {
		fmt.Println("Error while reading .json")
	}

	//convert jsondata to student slice
	studentlist := make([]user.Student, 0)
	err = json.Unmarshal(jsondata, &studentlist)
	if err != nil {
		fmt.Println(err)
	}

	//gebe Gruppenfarben
	for i, _ := range studentlist {
		switch {
		case i%3 == 0:
			studentlist[i].Gruppenfarbe = 1
		case i%3 == 1:
			studentlist[i].Gruppenfarbe = 2
		case i%3 == 2:
			studentlist[i].Gruppenfarbe = 3
		default:
			studentlist[i].Gruppenfarbe = 0
		}
	}

	tmpl, _ := template.ParseFiles("studentlist.go.html")
	tmpl.Execute(w, studentlist)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("register.go.html"))

	vorname := r.FormValue("vorname")
	nachname := r.FormValue("nachname")
	matrikel, _ := strconv.ParseInt(r.FormValue("matrikel")[0:], 10, 64)

	tpl.Execute(w, vorname)
	user.NewStudent(r.FormValue("vorname"), nachname, int(matrikel)).Register()
}

func main() {

	st1 := user.NewStudent("Kevin", "Kuhl", 212435)
	st1.Register()
	var user1 user.User
	user1 = user.NewStudent("Klaus", "Kruse", 111)
	user1.Unregister()
	//st1.Unregister()
	http.HandleFunc("/", handleRegister)
	http.ListenAndServe(":1313", nil)
	st2 := user.NewStudent("Kevin", "Kuhl", 350)
	st3 := user.NewStudent("Karsten", "Kerner", 435)
	user.Register(st2, st3)
}
