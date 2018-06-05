package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"./handlers"
	"github.com/gorilla/mux"
)

var conf handlers.Config = handlers.GetConfig()

func genUrl(str string) string {
	fmt.Printf("%s%s\n", conf.Root_url, str)
	return fmt.Sprintf("%s%s", conf.Root_url, str)
}

func main() {

	rtr := mux.NewRouter()

	FileServer := http.FileServer(http.Dir("static"))
	http.Handle(genUrl("static/"), http.StripPrefix(genUrl("static/"), FileServer))

	FileServerPortraits := http.FileServer(http.Dir("Userdata/Portraits"))
	http.Handle(genUrl("portraits/"), http.StripPrefix(genUrl("portraits/"), FileServerPortraits))

	rtr.HandleFunc(genUrl("register"), handlers.HandleRegister)
	rtr.HandleFunc(genUrl("login"), handlers.HandleLogin)
	rtr.HandleFunc(genUrl("post"), handlers.HandlePostForm)
	rtr.HandleFunc(genUrl("info"), handlers.HandleInfo)
	rtr.HandleFunc(genUrl("{number}"), handlers.HandleStudents)
	rtr.HandleFunc(genUrl(""), handlers.HandleStudents)
	rtr.HandleFunc(genUrl("profile/{matrikel}"), handlers.HandleProfile)
	rtr.HandleFunc(genUrl("{matrikel}/postraw/{postnr}"), handlers.HandleRawPosts)
	rtr.HandleFunc(genUrl("{matrikel}/post/{postnr}"), handlers.HandlePosts)
	http.Handle(genUrl(""), rtr)

	fmt.Printf("Start server at port%s:\n"+
		"   Coursname: %s\n"+
		"   Number of studentgroups: %v\n"+
		"   Course open: %v\n",
		conf.Port, conf.Course_name, conf.Group_number, conf.Open_course)

	//Start Server or exit with error message alter 5 seconds
	//err := http.ListenAndServeTLS(conf.Port, "../.tls/fullchain.pem", "../.tls/privkey.pem", nil)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second * 5)
		os.Exit(1)
	}
}
