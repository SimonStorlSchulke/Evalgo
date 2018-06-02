package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const MASTERPASSWORD string = "aeibom"

type Prof struct {
	Vorname  string `json:"vorname"`
	Nachname string `json:"nachname"`
}

func NewProf(Vorname, Nachname string) Prof {
	return Prof{Vorname, Nachname}
}

//Returns Folderpath to Profdata as string TODO: Besserer Pfadname (Momentan Nachname)
func (pr Prof) getPath() string {
	return filepath.Join(".", "Userdata", "Profs", fmt.Sprintf("%s", pr.Nachname))
}

func (pr Prof) ToJSON() []byte {
	jsondata, err := json.Marshal(pr)
	if err != nil {
		fmt.Println("Error when converting", pr, "to JSON")
		return nil
	}
	return jsondata
}

func (pr Prof) Unregister() {
	os.RemoveAll(pr.getPath())
	fmt.Printf("Unregistered Prof %s %s at %s\n", pr.Vorname, pr.Nachname, pr.getPath())
}

func (pr Prof) Register() {
	//create Subfolder in Studentdata/Matrikel
	fmt.Println(pr.getPath())
	err := os.MkdirAll(pr.getPath(), 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Write Profdata
	jsondata := pr.ToJSON()
	profilePath := filepath.Join(pr.getPath(), "profile.json")
	err = ioutil.WriteFile(profilePath, jsondata, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("registered Prof %s %s at %s\n", pr.Vorname, pr.Nachname, pr.getPath())

	//TODO: Check if already registered Matrikel
}

func (pr Prof) GetPassword() string {
	return MASTERPASSWORD
}
