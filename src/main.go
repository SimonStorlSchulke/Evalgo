package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Student struct {
	Name struct {
		Vorname  string `json:"vorname"`
		Nachname string `json:"nachname"`
	} `json:"name"`
	Matrikel int    `json:"matrikel"`
	Gruppe   int    `json:"gruppe"`
	Info     string `json:"info"`
	ID       int    `json:"id"`
}

//returns path to student-portrait
func (st *Student) PortraitPath() string {
	return fmt.Sprintf("/portraits/%v.jpg", st.Matrikel)
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

	tmpl, _ := template.ParseFiles("studentlist.go.html")
	tmpl.Execute(w, studentlist)
}

func main() {

	http.HandleFunc("/", handleStudents)
	http.ListenAndServe(":8080", nil)
}
