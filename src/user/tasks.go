package user

import (
	"fmt"
	"io/ioutil"
)

//Aus faulheit heraus im userpacket...
func GetTask(postNr int) []byte {
	nrStr, _ := intToString(postNr)
	task, err := ioutil.ReadFile(fmt.Sprintf("./coursedata/tasks/post_%s.md", nrStr))
	if err != nil {
		fmt.Println(err)
		return []byte("Noch keine Aufgabe.")
	}
	return task
}
