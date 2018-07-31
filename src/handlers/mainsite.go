package handlers

import (
	"fmt"
	"log"
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
		//TODO studenten auch Posts von authorisierte Nutzern anzeigen
	}

	//apply group colors TODO expose colors in coursecofig.json
	c1, c2, c3, cDef := "#beffa3", "#a3e6ff", "#f7ffa8", "#e2e2e2"
	grNum := conf.Group_number
	for i := range studentlist {
		switch {
		case i%grNum == 0:
			studentlist[i].Gruppenfarbe = c1
		case i%grNum == 1:
			studentlist[i].Gruppenfarbe = c2
		case i%grNum == 2:
			studentlist[i].Gruppenfarbe = c3
		default:
			studentlist[i].Gruppenfarbe = cDef
		}
	}

	tmpl, err := template.ParseFiles("./templates/mainsite.go.html")
	if err != nil {
		log.Fatalln(err)
	}

	var portraits []string
	for _, st := range studentlist {
		portraits = append(portraits, st.GetPortraitPath())
	}

	pageData := struct {
		Students        []user.User
		CurrentUser     user.User
		Portraits       []string
		CurrentPortrait string
		CourseName      string
		GradesEnabled   bool
		CardsEnabled    bool
	}{
		Students:        studentlist,
		CurrentUser:     currentUser,
		Portraits:       portraits,
		CurrentPortrait: currentUser.GetPortraitPath(),
		CourseName:      conf.Course_name,
		GradesEnabled:   conf.Enable_grades,
		CardsEnabled:    conf.Enable_cards,
	}

	//Feedback
	selectedUsStr, _ := r.URL.Query()["mat"]
	selectedPostStr, _ := r.URL.Query()["nr"]

	selectedUs, _ := strconv.Atoi(selectedUsStr[0])
	selectedPost, _ := strconv.Atoi(selectedPostStr[0])

	fbText := r.FormValue("fb-text")
	fbGrade, _ := strconv.Atoi(r.FormValue("fb-grade"))
	//fbCard, _ := strconv.Atoi(r.FormValue("feedback-text"))
	//fbCard := 0
	feedback := user.NewFeedback(fbText, fbGrade, 0)
	user.StoreFeedback(selectedUs, selectedPost, feedback)
	fmt.Println(feedback)

	tmpl.Execute(w, pageData)
}
