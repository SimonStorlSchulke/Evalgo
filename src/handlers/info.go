package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/russross/blackfriday"
)

//Course Info
func HandleInfo(w http.ResponseWriter, r *http.Request) {

	info, err := ioutil.ReadFile("./coursedata/info.md")
	if err != nil {
		fmt.Fprint(w, "no infofile found")
		return
	}

	//Parse Markdown via Blackfriday
	/*Blackfriday for some reasons always adds a Paragraph <p></p> around the first line... why though?
	Dirty fix is to start converting the info []byte at index 3 and therefore remove the first <p>*/
	info = blackfriday.MarkdownCommon(info[3:])
	str := "<div class='container-fluid post-area'>\n\n\n" + string(info[0:]) + "</div>"
	fmt.Fprint(w, str)
}
