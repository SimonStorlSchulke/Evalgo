package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

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

//check if user is authorized
func (us User) IsAuthorized() bool {
	return us.Usertype > STUDENT
}

//Determines whether a user may post or not
func (us User) MayPost() bool {
	//If User is student or all Tutors can post TODO nicht immer GetConfig aufrufen weil Performance
	return us.Usertype == STUDENT || courseconfig.GetConfig().Tutors_can_post
}

//Returns a User based on given Matrikel
func FromMatrikel(matrikel int) (User, error) {
	var us User
	jsondata, err := ioutil.ReadFile(fmt.Sprintf("./coursedata/users/%v/profile.json", matrikel))
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

//Return path to user portrait TODO: jpg support
func (us *User) GetPortraitPath() string {
	url := fmt.Sprintf("./portraits/%v.png", us.Matrikel)
	filepath := fmt.Sprintf("./coursedata/portraits/%v.png", us.Matrikel)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return "./portraits/default.png"
	}
	return url
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
	folders, err := ioutil.ReadDir("./coursedata/users")
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

/*ReadStudents reads all profile.json files in /coursedata/Students
and return it as a slice of Students*/
func ReadStudents() []User {
	//Read Folders
	folders, err := ioutil.ReadDir("./coursedata/users")
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
		path = fmt.Sprintf("./coursedata/users/%s/profile.json", file.Name())

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

/*ReadStudents reads all profile.json files in /coursedata/Students
and return it as a slice of Students*/
func ReadTutors() []User {
	//Read Folders
	folders, err := ioutil.ReadDir("./coursedata/users")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if !courseconfig.GetConfig().Tutors_can_post {
		return nil
	}

	var currentUser User
	var path string
	var jsondata []byte
	userlist := make([]User, 0)

	//Loop through Folders and create User Slice
	for _, file := range folders {
		path = fmt.Sprintf("./coursedata/users/%s/profile.json", file.Name())

		jsondata, err = ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		err = json.Unmarshal(jsondata, &currentUser)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if currentUser.Usertype > STUDENT {
			userlist = append(userlist, currentUser)
		}
	}
	return userlist
}

//---Util---

//Returns Folderpath to Studentdata as string
func (us *User) getPath() string {
	return filepath.Join(".", "coursedata", "users", fmt.Sprintf("%v", us.Matrikel))
}

//aplies user at beginning of Slice
func prepend(arr []User, item User) []User {
	return append([]User{item}, arr...)
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
