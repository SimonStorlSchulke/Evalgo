package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type Student struct {
	Vorname      string `json:"vorname"`
	Nachname     string `json:"nachname"`
	Matrikel     int    `json:"matrikel"`
	Gruppenfarbe string
	Info         string `json:"info"`
	ID           int    `json:"id"`
}

//Constructor for Student only Vorname, Nachname, Matrikel
func NewStudent(Vorname, Nachname string, Matrikel int) Student {
	return Student{Vorname, Nachname, Matrikel, "", "", 0}
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

//Returns Posts of a Student
func (st Student) GetPosts() []byte {

	//TODO: das gleiche mit input matrikel

	posts, err := ioutil.ReadFile(fmt.Sprintf("./Userdata/Students/%v/posts.md", st.Matrikel))
	if err != nil {
		fmt.Println(err)
		return []byte("Noch keine Einträge.")
	}
	return posts
}

func (st Student) ToJSON() []byte {
	jsondata, err := json.Marshal(st)
	if err != nil {
		fmt.Println("Error when converting", st, "to JSON")
		return nil
	}
	return jsondata
}

//Write Student JSON to Studentdata Folder
func (st Student) Register() {

	//return if Matrikel already exists
	for _, existingSt := range ReadStudents() {
		if existingSt.Matrikel == st.Matrikel {
			fmt.Println("Student with the Matrikel", st.Matrikel, "already exists")
			return
		}
	}

	//return if Matrikel = 0
	if st.Matrikel == 0 {
		fmt.Println("No Matrikelnumber given")
		return
	}

	//create Subfolder in Studentdata/Matrikel
	fmt.Println(st.getPath())
	err := os.MkdirAll(st.getPath(), 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Write Studentdata to Studentdata/Matrikel/profile.json
	jsondata := st.ToJSON()
	profilePath := filepath.Join(st.getPath(), "profile.json")
	err = ioutil.WriteFile(profilePath, jsondata, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("registered Student %s %s at %s\n", st.Vorname, st.Nachname, st.getPath())
	//TODO: Check if already registered Matrikel!!
}

//Removes Studentdata from /Studentdata
//TODO - macht das sinn als Methode? Oder lieber als Funktion?
func (st Student) Unregister() {
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

//Returns Folderpath to Studentdata as string
func (st Student) getPath() string {
	return filepath.Join(".", "Userdata", "Students", fmt.Sprintf("%v", st.Matrikel))
}
