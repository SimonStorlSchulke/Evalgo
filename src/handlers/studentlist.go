package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"../user"
)

type Config struct {
	Port         string `json:"port"`
	Course_name  string `json:"course_name"`
	Open_course  bool   `json:"open_course"`
	Group_number int    `json:"group_number"`
	Root_url     string `json:"root_url"`
}

//Read courseconfig.json and return as Config struct
func GetConfig() Config {
	var conf Config
	jsondata, err := ioutil.ReadFile("./courseconfig.json")
	if err != nil {
		fmt.Println(err)
		return Config{Open_course: false}
	}

	err = json.Unmarshal(jsondata, &conf)
	if err != nil {
		fmt.Println(err)
		return Config{Open_course: false}
	}
	return conf
}

//Handles Main Site
func HandleStudents(w http.ResponseWriter, r *http.Request) {
	conf := GetConfig()
	var currentUser user.Student
	var err error
	loggedIn, mat := loggedIn(r)
	if loggedIn {
		currentUser, err = user.FromMatrikel(mat)
	}

	//Read jsons:
	studentlist := []user.Student{}
	if conf.Open_course {
		studentlist = user.ReadStudents()
	} else {
		studentlist = []user.Student{currentUser}
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

	tmpl, err := template.ParseFiles("./templates/studentlist.go.html")
	if err != nil {
		log.Fatalln(err)
	}

	var portraits []string
	for _, st := range studentlist {
		portraits = append(portraits, st.GetPortraitPath())
	}

	pageData := struct {
		Nav             string
		Students        []user.Student
		CurrentUser     user.Student
		Portraits       []string
		CurrentPortrait string
	}{
		Nav:             getNav(),
		Students:        studentlist,
		CurrentUser:     currentUser,
		Portraits:       portraits,
		CurrentPortrait: currentUser.GetPortraitPath(),
	}

	tmpl.Execute(w, pageData)
}
