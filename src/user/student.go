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
)

type Student struct {
	Vorname      string `json:"vorname"`
	Nachname     string `json:"nachname"`
	Matrikel     int    `json:"matrikel"`
	Gruppenfarbe string
	Info         string `json:"info"`
	Passwort     string `json:"passwort"`
}

//Constructor for Student only Vorname, Nachname, Matrikel
func NewStudent(Vorname, Nachname string, Matrikel int, Passwort string) Student {
	return Student{Vorname, Nachname, Matrikel, "", "", Passwort}
}

//Returns a Student based on given Matrikel
func FromMatrikel(matrikel int) (Student, error) {
	var st Student
	jsondata, err := ioutil.ReadFile(fmt.Sprintf("./Userdata/Students/%v/profile.json", matrikel))
	if err != nil {
		fmt.Println(err)
		return st, err
	}

	err = json.Unmarshal(jsondata, &st)
	if err != nil {
		fmt.Println(err)
	}
	return st, err
}

//Post a string to /Userdata/Students/[st.matrikel]/post_[postnumber].md
func (st *Student) PostNr(str string, postNumber int) {
	ioutil.WriteFile(fmt.Sprintf("./Userdata/Students/%v/post_%v.md", st.Matrikel, postNumber), []byte(str), 0777)
	fmt.Println(st.Vorname, st.Nachname, "created a new post Nr.", postNumber)
}

func (st *Student) GetPost(postNr int) []byte {

	post, err := ioutil.ReadFile(fmt.Sprintf("./Userdata/Students/%v/post_%v.md", st.Matrikel, postNr))
	if err != nil {
		return []byte("Noch keine Abgabe.")
	}
	return post
}

//Returns []byte of all posts and int slice of postNumbers
func (st *Student) GetAllPosts() ([]byte, []int) {
	posts, _ := ioutil.ReadDir(fmt.Sprintf("./Userdata/Students/%v/", st.Matrikel))
	var postdata []byte
	var data []byte
	var postNumbers []int
	for _, p := range posts {
		number, err := strconv.Atoi(strings.Trim(p.Name(), "post_.md"))
		if err == nil {
			currentPostStr := fmt.Sprintf("# <div class='post-header text-primary'>Aufgabe %v</div>\n", number)
			currentPost := []byte(currentPostStr)
			postdata = st.GetPost(number)
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
func (st *Student) GetPortraitPath() string {
	url := fmt.Sprintf("./portraits/%v.png", st.Matrikel)
	filepath := fmt.Sprintf("./Userdata/Portraits/%v.png", st.Matrikel)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return "./portraits/default.png"
	}
	return url
}

//Returns Folderpath to Studentdata as string
func (st *Student) getPath() string {
	return filepath.Join(".", "Userdata", "Students", fmt.Sprintf("%v", st.Matrikel))
}

//Returns Password of Student
func (st *Student) GetPassword() string {
	return st.Passwort
}

//Convert Student Struct to JSON
func (st *Student) ToJSON() []byte {
	jsondata, err := json.Marshal(st)
	if err != nil {
		fmt.Println("Error when converting", st, "to JSON")
		return nil
	}
	return jsondata
}

//Write Student JSON to Studentdata Folder
func (st Student) Register() error {

	//return if Matrikel already exists
	for _, existingSt := range ReadStudents() {
		if existingSt.Matrikel == st.Matrikel {
			return errors.New("Student with the Matrikel already exists")
		}
	}

	//return if Matrikel = 0
	if st.Matrikel == 0 {
		return errors.New("No Matrikelnumber given")
	}

	//create Subfolder in Studentdata/Matrikel
	err := os.MkdirAll(st.getPath(), 0777)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//Write Studentdata to Studentdata/Matrikel/profile.json
	jsondata := st.ToJSON()
	profilePath := filepath.Join(st.getPath(), "profile.json")
	err = ioutil.WriteFile(profilePath, jsondata, 0777)
	if err != nil {
		return err
	}

	fmt.Printf("registered Student %s %s at %s\n", st.Vorname, st.Nachname, st.getPath())
	return nil
}

//Removes Studentdata from /Studentdata
//TODO: - expose to user
func (st *Student) Unregister() {
	os.RemoveAll(st.getPath())
	fmt.Printf("Unregistered Student %s %s at %s\n", st.Vorname, st.Nachname, st.getPath())
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
func ReadStudents() []Student {
	//Read Folders
	folders, err := ioutil.ReadDir("./Userdata/Students")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var currentStudent Student
	var path string
	var jsondata []byte
	studentlist := make([]Student, 0)

	//Loop through Folders and create Student Slice
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

		studentlist = append(studentlist, currentStudent)
	}
	return studentlist
}
