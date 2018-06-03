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

	info = blackfriday.MarkdownCommon(info)
	//Parse Markdown to []byte
	fmt.Fprint(w, string(info[:]))
}
