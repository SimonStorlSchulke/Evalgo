package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"

	"../courseconfig"
	"../user"
)

//Handles Main Site
func HandleMainSite(w http.ResponseWriter, r *http.Request) {
	conf := courseconfig.GetConfig()
	var currentUser user.User
	var err error
	loggedIn, mat := loggedIn(r)
	if loggedIn {
		currentUser, err = user.FromMatrikel(mat)
	}

	//Read jsons:
	studentlist := []user.User{}
	if conf.Open_course || currentUser.IsAuthorized() {
		studentlist = user.ReadStudents()
	} else {
		studentlist = user.ReadTutors()
		studentlist = append(studentlist, currentUser)
	}

	//Extract Students from Users and combine them again later - there must be a better way though...
	var onlyStudents []user.User
	var otherUsers []user.User

	for _, us := range studentlist {
		if us.Usertype == user.STUDENT {
			onlyStudents = append(onlyStudents, us)
		} else {
			otherUsers = append(otherUsers, us)
		}
	}

	/*append groupcolor classes (actual colors are defined in groupcolors.css)
	TODO: Display Groupnumbers correctly in closed course.*/
	groupColors := []string{"grc-1", "grc-2", "grc-3", "grc-4", "grc-5", "grc-6", "grc-7", "grc-8", "grc-9", "grc-10", "grc-11"}
	cNum := conf.Group_number
	//Limit maxG Group Number to 6 - maybe revisit this later
	if cNum > 6 {
		cNum = 6
	}
	for i := range onlyStudents {
		v := (float32(i) / float32(len(onlyStudents))) * float32(cNum)
		onlyStudents[i].Gruppenfarbe = groupColors[int(v)]
	}

	studentlist = append(otherUsers, onlyStudents...)

	tmpl, err := template.ParseFiles("./templates/mainsite.go.html")
	if err != nil {
		log.Fatalln(err)
	}

	var portraits []string
	for _, st := range studentlist {
		portraits = append(portraits, st.GetPortraitPath())
	}

	//random number to be added at the end of profile pic links to force the browser to reload them each time

	pageData := struct {
		Students        []user.User
		CurrentUser     user.User
		Portraits       []string
		CurrentPortrait string
		CourseName      string
		Conf            courseconfig.Config
		TaskNumbers     []int
		Rand            int
	}{
		Students:        studentlist,
		CurrentUser:     currentUser,
		Portraits:       portraits,
		CurrentPortrait: currentUser.GetPortraitPath(),
		CourseName:      conf.Course_name,
		Conf:            conf,
		TaskNumbers:     existingTaskNumbers(),

		//random number to be added at the end of profile pic links to force the browser to reload them each time (and not use the old one from cache)
		Rand: rand.Int(),
	}

	//Feedback
	selectedUsStr, paramsOk := r.URL.Query()["mat"]
	selectedPostStr, paramsOk := r.URL.Query()["nr"]

	if paramsOk {
		selectedUs, _ := strconv.Atoi(selectedUsStr[0])
		selectedPost, _ := strconv.Atoi(selectedPostStr[0])

		if selectedUs > 0 && selectedPost > 0 {

			fbText := r.FormValue("fb-text")
			fbGrade, _ := strconv.Atoi(r.FormValue("fb-grade"))
			fbCard, _ := strconv.Atoi(r.FormValue("fb-card"))

			//Store Feedback
			if fbGrade != 0 || !conf.Enable_grades && fbText != "" {
				feedback := user.NewFeedback(fbText, fbGrade, fbCard)
				err = user.StoreFeedback(selectedUs, selectedPost, feedback)
				if err != nil {
					fmt.Println(err)
				}
			}

		}
	}

	tmpl.Execute(w, pageData)
}
