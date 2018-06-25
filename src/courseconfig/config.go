package courseconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Port            string `json:"port"`
	Course_name     string `json:"course_name"`
	Open_course     bool   `json:"open_course"`
	Group_number    int    `json:"group_number"`
	Root_url        string `json:"root_url"`
	Master_password string `json:"master_password"`
	Tutors_can_post bool   `json:"tutors_can_post"`
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
