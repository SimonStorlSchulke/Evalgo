package handlers

import (
	"net/http"
	"text/template"

	"../user"
)

func HandleTable(w http.ResponseWriter, r *http.Request) {

	tasks := existingTaskNumbers()
	students := user.ReadStudents()
	var studentsOnly []user.User
	for _, st := range students {
		if st.Usertype == user.STUDENT {
			studentsOnly = append(studentsOnly, st)
		}
	}

	//Keys: Matrikel, Tasknumber
	grades := make(map[int]map[int]int)

	for _, st := range studentsOnly {
		grades[st.Matrikel] = make(map[int]int)
		for _, ts := range tasks {
			grade := 0
			fb, err := user.GetFeedback(st.Matrikel, ts)
			if err == nil {
				grade = fb.Grade
			}
			grades[st.Matrikel][ts] = grade
		}

	}

	pageData := struct {
		Tasks    []int
		Students []user.User
		Grades   map[int]map[int]int
	}{tasks, studentsOnly, grades}

	tpl := template.Must(template.ParseFiles("./templates/table.go.html"))
	tpl.Execute(w, pageData)
}
