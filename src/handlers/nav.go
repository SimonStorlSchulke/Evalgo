package handlers

import "io/ioutil"

func getNav() string {
	headerData, _ := ioutil.ReadFile("./templates/nav.html")
	return string(headerData[:])
}
