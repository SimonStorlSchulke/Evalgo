package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type htmlcolor string

const (
	rot  = "#ff8080"
	grün = "#66ff99"
	blau = "#33ccff"
	grau = "#969696"
)

type Student struct {
	Name struct {
		Vorname  string `json:"vorname"`
		Nachname string `json:"nachname"`
	} `json:"name"`
	Matrikel     int `json:"matrikel"`
	Gruppenfarbe htmlcolor
	Info         string `json:"info"`
	ID           int    `json:"id"`
}

//returns path to student-portrait
func (st *Student) PortraitPath() string {
	return fmt.Sprintf("/portraits/%v.jpg", st.Matrikel)
}

//momentan unbenutzt - funzt iwie ned (farben werden nicht angewendet)
func verteileGruppen(stl *[]Student) {
	for i, stu := range *stl {
		switch {
		case i%3 == 0:
			stu.Gruppenfarbe = rot
		case i%3 == 1:
			stu.Gruppenfarbe = grün
		case i%3 == 2:
			stu.Gruppenfarbe = blau
		default:
			stu.Gruppenfarbe = grau
		}
	}
}

func handleStudents(w http.ResponseWriter, r *http.Request) {
	//Read json:
	jsondata, err := ioutil.ReadFile("./students.json")
	if err != nil {
		fmt.Println("Error while reading .json")
	}

	//convert jsondata to student slice
	studentlist := make([]Student, 0)
	err = json.Unmarshal(jsondata, &studentlist)
	if err != nil {
		fmt.Println("Error while converting jsondata")
	}

	//gebe Gruppenfarben
	for i, _ := range studentlist {
		switch {
		case i%3 == 0:
			studentlist[i].Gruppenfarbe = rot
		case i%3 == 1:
			studentlist[i].Gruppenfarbe = grün
		case i%3 == 2:
			studentlist[i].Gruppenfarbe = blau
		default:
			studentlist[i].Gruppenfarbe = grau
		}
	}

	tmpl, _ := template.ParseFiles("studentlist.go.html")
	tmpl.Execute(w, studentlist)
}

func main() {

	http.HandleFunc("/", handleStudents)
	http.ListenAndServe(":8080", nil)
}
