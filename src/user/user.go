package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"../courseconfig"
)

type User struct {
	Vorname      string `json:"vorname"`
	Nachname     string `json:"nachname"`
	Matrikel     int    `json:"matrikel"`
	Gruppenfarbe string
	Passwort     string   `json:"passwort"`
	Usertype     Usertype `json:"usertype"`
}

// Defines the type of User
type Usertype int

const (
	/*STUDENT has the rights to post assignments*/
	STUDENT Usertype = iota + 1
	/*TUTOR has the right to rate Assignments and give Feedback*/
	TUTOR
	/*ADMIN has the right to rate to create Assignments,
	Assignments and give Feedback, unregister students,*/
	ADMIN
)

//Constructor for User only Vorname, Nachname, Matrikel
func NewUser(Vorname, Nachname string, Matrikel int, Passwort string) User {
	return User{Vorname, Nachname, Matrikel, "", Passwort, STUDENT}
}

//Constructor for User only Vorname, Nachname, Matrikel
func NewAuthUser(Vorname, Nachname string, Matrikel int, Passwort string) User {
	return User{Vorname, Nachname, Matrikel, "", Passwort, TUTOR}
}

//Returns true if User is of Usertype > Student
func (us User) IsAuthorized() bool {
	if us.Usertype > STUDENT {
		return true
	}
	return false
}

//Determines whether a user may post or not
func (us User) MayPost() bool {
	//If User is student or all Tutors can post TODO nicht immer GetConfig aufrufen weil Performance
	if us.Usertype == STUDENT || courseconfig.GetConfig().Tutors_can_post {
		return true
	}
	return false
}

//Returns a User based on given Matrikel
func FromMatrikel(matrikel int) (User, error) {
	var us User
	jsondata, err := ioutil.ReadFile(fmt.Sprintf("./Userdata/Students/%v/profile.json", matrikel))
	if err != nil {
		fmt.Println(err)
		return us, err
	}

	err = json.Unmarshal(jsondata, &us)
	if err != nil {
		fmt.Println(err)
	}
	return us, err
}

//Converts an int to a three digit string (for example 5 -> "005")
func intToString(num int) (string, error) {
	if num > 999 || num < 0 {
		return "", errors.New("Number must be < than 1000 and > 0")
	}
	str := strconv.Itoa(num)
	for mDg := 3 - len(str); mDg > 0; mDg-- {
		str = "0" + str
	}
	return str, nil
}

//Post a string to /Userdata/Students/[st.matrikel]/post_[postnumber].md
func (us *User) PostNr(str string, postNumber int) {
	nrStr, err := intToString(postNumber)
	if err == nil {
		ioutil.WriteFile(fmt.Sprintf("./Userdata/Students/%v/post_%s.md", us.Matrikel, nrStr), []byte(str), 0777)
		fmt.Println(us.Vorname, us.Nachname, "created a new post Nr.", postNumber)
	}
	//TODO except
}

//Returns a post as []byte
func (us *User) GetPost(postNr int) []byte {
	nrStr, _ := intToString(postNr)
	post, err := ioutil.ReadFile(fmt.Sprintf("./Userdata/Students/%v/post_%s.md", us.Matrikel, nrStr))
	if err != nil {
		return []byte("Noch keine Abgabe.")
	}
	return post
}

//Returns []byte containing all posts of student and []int of postNumbers
func (us *User) GetAllPosts() ([]byte, []int) {
	posts, _ := ioutil.ReadDir(fmt.Sprintf("./Userdata/Students/%v/", us.Matrikel))
	var postdata []byte
	var data []byte
	var postNumbers []int
	for _, p := range posts {
		number, err := strconv.Atoi(strings.Trim(p.Name(), "post_.md"))
		if err == nil {
			currentPostStr := fmt.Sprintf("# <div class='post-header text-primary'>Aufgabe %v</div>\n", number)
			currentPost := []byte(currentPostStr)
			postdata = us.GetPost(number)
			currentPost = append(currentPost, postdata...)
			hr := []byte("\n\n<hr>\n")
			currentPost = append(currentPost, hr...)
			data = append(data, currentPost...)
			postNumbers = append(postNumbers, number)
		}
	}
	return data, postNumbers
}

//Return path to user portrait TODO: jpg.
func (us *User) GetPortraitPath() string {
	url := fmt.Sprintf("./portraits/%v.png", us.Matrikel)
	filepath := fmt.Sprintf("./Userdata/Portraits/%v.png", us.Matrikel)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return "./portraits/default.png"
	}
	return url
}

//Returns Folderpath to Studentdata as string
func (us *User) getPath() string {
	return filepath.Join(".", "Userdata", "Students", fmt.Sprintf("%v", us.Matrikel))
}

//Returns Password of User
func (us *User) GetPassword() string {
	return us.Passwort
}

//Convert User Struct to JSON
func (us *User) ToJSON() []byte {
	jsondata, err := json.Marshal(us)
	if err != nil {
		fmt.Println("Error when converting", us, "to JSON")
		return nil
	}
	return jsondata
}

//Write User JSON to Studentdata Folder
func (us User) Register() error {

	//return if Matrikel already exists
	for _, existingSt := range ReadStudents() {
		if existingSt.Matrikel == us.Matrikel {
			return errors.New("User with the Matrikel already exists")
		}
	}

	//return if Matrikel = 0
	if us.Matrikel == 0 {
		return errors.New("No Matrikelnumber given")
	}

	//create Subfolder in Studentdata/Matrikel
	err := os.MkdirAll(us.getPath(), 0777)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//Write Studentdata to Studentdata/Matrikel/profile.json
	jsondata := us.ToJSON()
	profilePath := filepath.Join(us.getPath(), "profile.json")
	err = ioutil.WriteFile(profilePath, jsondata, 0777)
	if err != nil {
		return err
	}

	fmt.Printf("registered User %s %s at %s\n", us.Vorname, us.Nachname, us.getPath())
	return nil
}

//Removes Studentdata from /Studentdata
//TODO: - expose to user
func (us *User) Unregister() {
	os.RemoveAll(us.getPath())
	fmt.Printf("Unregistered User %s %s at %s\n", us.Vorname, us.Nachname, us.getPath())
}

//returns a List of registered Matrikelnumbers
func MatrikelList() []int {
	folders, err := ioutil.ReadDir("./Userdata/Students")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	mList := make([]int, 0)
	for _, file := range folders {
		m, err := strconv.Atoi(file.Name())
		if err != nil {
			fmt.Println(err)
			return nil
		}
		mList = append(mList, m)
	}
	return mList
}

/*ReadStudents reads all profile.json files in /Userdata/Students
and return it as a slice of Students*/
func ReadStudents() []User {
	//Read Folders
	folders, err := ioutil.ReadDir("./Userdata/Students")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var currentStudent User
	var path string
	var jsondata []byte
	studentlist := make([]User, 0)

	//Loop through Folders and create User Slice
	for _, file := range folders {
		//TODO user FromMatrikel here
		path = fmt.Sprintf("./Userdata/Students/%s/profile.json", file.Name())

		jsondata, err = ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		err = json.Unmarshal(jsondata, &currentStudent)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if currentStudent.Usertype == STUDENT {
			studentlist = append(studentlist, currentStudent)
		} else if courseconfig.GetConfig().Tutors_can_post {
			studentlist = prepend(studentlist, currentStudent)
		}
	}
	return studentlist
}

//aplies user at beginning of Slice
func prepend(arr []User, item User) []User {
	return append([]User{item}, arr...)
}
