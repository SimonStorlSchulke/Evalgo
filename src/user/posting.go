package user

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

//Post a string to /coursedata/users/[st.matrikel]/post_[postnumber].md
func (us *User) PostNr(str string, postNumber int) error {
	nrStr, err := intToString(postNumber)
	if err != nil {
		return err
	}
	//Add Time Stamp
	tm := time.Now()
	timestring := fmt.Sprintf("<i class='last-modified'>last modified %v:%v %s %v %v</i>\n", tm.Hour(), tm.Minute(), tm.Month(), tm.Day(), tm.Year())
	str = timestring + str

	err = ioutil.WriteFile(fmt.Sprintf("./coursedata/users/%v/post_%s.md", us.Matrikel, nrStr), []byte(str), 0777)
	if err != nil {
		return err
	}
	fmt.Println(us.Vorname, us.Nachname, "created post Nr.", postNumber)
	return nil
}

//Returns a post as []byte
func (us *User) GetPost(postNr int) []byte {
	nrStr, _ := intToString(postNr)
	post, err := ioutil.ReadFile(fmt.Sprintf("./coursedata/users/%v/post_%s.md", us.Matrikel, nrStr))
	if err != nil {
		return []byte("Noch keine Abgabe.")
	}
	return post
}

//Returns []byte containing all posts of student and []int of postNumbers
func (us *User) GetAllPosts() ([]byte, []int) {
	posts, _ := ioutil.ReadDir(fmt.Sprintf("./coursedata/users/%v/", us.Matrikel))
	var postdata []byte
	var data []byte
	var postNumbers []int
	for _, p := range posts {
		number, err := strconv.Atoi(strings.Trim(p.Name(), "post_.md"))
		if err == nil {
			currentPostStr := fmt.Sprintf("# <div class='post-header text-primary'><a href='./?nr=%v&mat=%v'>Aufgabe %v<a/></div>\n", number, us.Matrikel, number)
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
