package handlers

import (
	"net/http"
	"text/template"

	"../user"
)

func HandleTable(w http.ResponseWriter, r *http.Request) {

	tasks := existingTaskNumbers()
	students := user.ReadStudents()

	pageData := struct {
		Tasks    []int
		Students []user.User
	}{tasks, students}

	tpl := template.Must(template.ParseFiles("./templates/table.go.html"))
	tpl.Execute(w, pageData)
}
