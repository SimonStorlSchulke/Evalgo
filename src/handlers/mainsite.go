package handlers

import (
	"log"
	"net/http"
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
	if conf.Open_course {
		studentlist = user.ReadStudents()
	} else {
		studentlist = []user.User{currentUser}
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
		Nav             string
		Students        []user.User
		CurrentUser     user.User
		Portraits       []string
		CurrentPortrait string
		CourseName      string
	}{
		Nav:             getNav(),
		Students:        studentlist,
		CurrentUser:     currentUser,
		Portraits:       portraits,
		CurrentPortrait: currentUser.GetPortraitPath(),
		CourseName:      conf.Course_name,
	}

	tmpl.Execute(w, pageData)
}
