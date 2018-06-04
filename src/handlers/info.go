package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/russross/blackfriday"
)

//Course Info
func HandleInfo(w http.ResponseWriter, r *http.Request) {

	info, err := ioutil.ReadFile("./info.md")
	if err != nil {
		fmt.Fprint(w, "no infofile found")
		return
	}

	//Parse Markdown and parse to string
	info = blackfriday.MarkdownCommon(info)
	fmt.Fprint(w, string(info[:]))
}
