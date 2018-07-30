package user

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//Post a string to /Userdata/Students/[st.matrikel]/post_[postnumber].md
func (us *User) PostNr(str string, postNumber int) {
	nrStr, err := intToString(postNumber)
	if err == nil {
		ioutil.WriteFile(fmt.Sprintf("./Userdata/Students/%v/post_%s.md", us.Matrikel, nrStr), []byte(str), 0777)
		fmt.Println(us.Vorname, us.Nachname, "created a new post Nr.", postNumber)
	}
}

//Returns a post as []byte
func (us *User) GetPost(postNr int) []byte {
	nrStr, _ := intToString(postNr)
	post, err := ioutil.ReadFile(fmt.Sprintf("./Userdata/Students/%v/post_%s.md", us.Matrikel, nrStr))
	if err != nil {
		return []byte("Noch keine Abgabe.")
	}
	return post
}

//Returns []byte containing all posts of student and []int of postNumbers
func (us *User) GetAllPosts() ([]byte, []int) {
	posts, _ := ioutil.ReadDir(fmt.Sprintf("./Userdata/Students/%v/", us.Matrikel))
	var postdata []byte
	var data []byte
	var postNumbers []int
	for _, p := range posts {
		number, err := strconv.Atoi(strings.Trim(p.Name(), "post_.md"))
		if err == nil {
			//TODO: clean up hard-coded html
			currentPostStr := fmt.Sprintf("# <div class='post-header text-primary'>Aufgabe %v</div>\n", number)
			currentPost := []byte(currentPostStr)
			postdata = us.GetPost(number)
			currentPost = append(currentPost, postdata...)
			hr := []byte("\n\n<hr>\n")
			currentPost = append(currentPost, hr...)
			data = append(data, currentPost...)
			postNumbers = append(postNumbers, number)
		}
	}
	return data, postNumbers
}
