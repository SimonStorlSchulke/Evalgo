package handlers

import (
	"net/http"
	"text/template"

	"../courseconfig"
	"../user"
)

func HandleTable(w http.ResponseWriter, r *http.Request) {

	if !isAuthSession(r, w) {
		WriteMsg(w, MsgPermissionDenied)
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

	fbt := FeedbackTable(&studentsOnly)

	pageData := struct {
		Coursename string
		Tasks      []int
		Students   []user.User
		Grades     map[int]map[int]user.Feedback
	}{cfg.Course_name, tasks, studentsOnly, fbt}

	tpl := template.Must(template.ParseFiles("./templates/table.go.html"))
	tpl.Execute(w, pageData)
}

//2D FeedbackMap with grades and cards of Students and Tasks
func FeedbackTable(students *[]user.User) map[int]map[int]user.Feedback {
	cfg := courseconfig.Conf
	tasks := existingTaskNumbers()

	//2D Feedback Map - Keys: Matrikel, Tasknumber
	fbt := make(map[int]map[int]user.Feedback)

	for _, st := range *students {
		fbt[st.Matrikel] = make(map[int]user.Feedback)
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
			fbt[st.Matrikel][ts] = fb
		}
	}
	return fbt
}
