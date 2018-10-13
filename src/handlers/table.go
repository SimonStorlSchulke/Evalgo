package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"../courseconfig"
	"../user"
)

func HandleTable(w http.ResponseWriter, r *http.Request) {

	//Check Permission and deny if no loggged in user or user is not authorized
	loggedIn, loggedInMat := loggedIn(r)
	loggedInUser, err := user.FromMatrikel(loggedInMat)

	if !loggedIn || loggedInUser.Usertype == user.STUDENT || err != nil {
		fmt.Fprintf(w, "Permission Denied. This Page is only accessible by authorized, logged in users.")
		return
	}

	cfg := courseconfig.Conf

	tasks := existingTaskNumbers()
	students := user.ReadStudents()

	//Filter out Tutors from Table
	var studentsOnly []user.User
	for _, st := range students {
		if st.Usertype == user.STUDENT {
			studentsOnly = append(studentsOnly, st)
		}
	}

	//2D Feedback Map - Keys: Matrikel, Tasknumber
	fbs := make(map[int]map[int]user.Feedback)

	for _, st := range studentsOnly {
		fbs[st.Matrikel] = make(map[int]user.Feedback)
		for _, ts := range tasks {
			fb := user.Feedback{}
			cfb, err := user.GetFeedback(st.Matrikel, ts)
			if err == nil {
				fb = cfb
				if !cfg.Enable_grades {
					fb.Grade = 0
				}
				if !cfg.Enable_cards {
					fb.Card = 0
				}
			}
			fbs[st.Matrikel][ts] = fb
		}

	}

	pageData := struct {
		Coursename string
		Tasks      []int
		Students   []user.User
		Grades     map[int]map[int]user.Feedback
	}{cfg.Course_name, tasks, studentsOnly, fbs}

	tpl := template.Must(template.ParseFiles("./templates/table.go.html"))
	tpl.Execute(w, pageData)
}
