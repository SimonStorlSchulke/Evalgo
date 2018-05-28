package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

func (st Student) ToJSON() []byte {
	jsondata, err := json.Marshal(st)
	if err != nil {
		fmt.Println("Error when converting", st, "to JSON")
		return nil
	}
	return jsondata
}

//Returns Folderpath to Studentdata as string
func (st Student) getPath() string {
	return filepath.Join(".", "Userdata", "Students", fmt.Sprintf("%v", st.Matrikel))
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

/*ReadStudents reads all profile.json files in /Userdata/Students
and return it as a slice of Students*/
func ReadStudents() []Student {
	//Read Folders
	folders, err := ioutil.ReadDir("./Userdata/Students")
	if err != nil {
		log.Fatal(err)
	}

	var currentStudent Student
	var path string
	var jsondata []byte
	studentlist := make([]Student, 0)

	//Loop through Folders and create Student Slice
	for _, file := range folders {
		path = fmt.Sprintf("./Userdata/Students/%s/profile.json", file.Name())
		jsondata, err = ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(jsondata, &currentStudent)
		if err != nil {
			log.Fatal(err)
		}
		studentlist = append(studentlist, currentStudent)
	}
	return studentlist
}
