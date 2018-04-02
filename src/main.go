package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Student struct {
	Vorname  string
	Name     string
	Matrikel int
	Gruppe   int
	Info     string
	ID       int
}

//Schreibt Daten des Studenten in Konsole
func (st *Student) PrintData() {
	fmt.Printf("%s %s:\n"+
		"Matrikel: \t %v\n"+
		"Gruppe: \t %v\n"+
		"Infofeld: \t %s\n\n",
		st.Vorname, st.Name, st.Matrikel, st.Gruppe, st.Info)
}

//gibt pfad zum Partrait aus
func (st *Student) getPortraitPath() string {
	return fmt.Sprintf("/portraits/%v.jpg", st.Matrikel)
}

func handleFunc(w http.ResponseWriter, r *http.Request) {

	studentList := []Student{
		Student{"Max", "Mustermann", 261812, 1, "Ich bin kuhl", 3},
		Student{"Kalle", "Klößchen", 214679, 2, "Ich bin kuhl", 3},
		Student{"Freddy", "Ferrari", 231678, 1, "Ich bin kuhl", 3},
		Student{"Vladimir", "Vlidimirovic", 296837, 2, "Ich bin kuhl", 3}}

	//parse template to html
	tmpl, _ := template.ParseFiles("studentlist.gohtml")
	//handle
	tmpl.Execute(w, studentList)

}

func main() {

	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":8080", nil)
}
