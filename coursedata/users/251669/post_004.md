<i class='last-modified'>last modified 10:25 October 15 2018</i>
## Read info.md File
```go
func HandleInfo(w http.ResponseWriter, r *http.Request) {

	info, err := ioutil.ReadFile("./info.md")
	if err != nil {
		fmt.Fprint(w, "no infofile found")
		return
	}

	//Parse Markdown and parse to string
	info = blackfriday.MarkdownCommon(info)
	str := "<div class='container-fluid post-area'>" + string(info[:]) + "</div>"
	fmt.Fprint(w, str)
}
```