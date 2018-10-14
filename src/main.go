package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"./courseconfig"
	"./handlers"
	"github.com/gorilla/mux"
)

var conf courseconfig.Config = courseconfig.GetConfig()

func genUrl(str string) string {
	fmt.Printf("%s%s\n", conf.Root_url, str)
	return fmt.Sprintf("%s%s", conf.Root_url, str)
}

func main() {

	rtr := mux.NewRouter()

	FileServer := http.FileServer(http.Dir("static"))
	http.Handle(genUrl("static/"), http.StripPrefix(genUrl("static/"), FileServer))

	FileServerPortraits := http.FileServer(http.Dir("coursedata/portraits"))
	http.Handle(genUrl("portraits/"), http.StripPrefix(genUrl("portraits/"), FileServerPortraits))

	FileServerImg := http.FileServer(http.Dir("coursedata/tasks/img"))
	http.Handle(genUrl("img/"), http.StripPrefix(genUrl("img/"), FileServerImg))

	rtr.HandleFunc(genUrl("register"), handlers.HandleRegister)
	rtr.HandleFunc(genUrl("login"), handlers.HandleLogin)
	rtr.HandleFunc(genUrl("authlogin"), handlers.HandleAuthLogin)
	rtr.HandleFunc(genUrl("portrait"), handlers.PortraitUpload)
	rtr.HandleFunc(genUrl("post"), handlers.HandlePostForm)
	rtr.HandleFunc(genUrl("table"), handlers.HandleTable)
	rtr.HandleFunc(genUrl("info"), handlers.HandleInfo)
	rtr.HandleFunc(genUrl("{number}"), handlers.HandleMainSite)
	rtr.HandleFunc(genUrl(""), handlers.HandleMainSite)
	rtr.HandleFunc(genUrl("profile/{matrikel}"), handlers.HandleProfile)
	rtr.HandleFunc(genUrl("{matrikel}/postraw/{postnr}"), handlers.HandleRawPosts)
	rtr.HandleFunc(genUrl("{matrikel}/pdf"), handlers.HandlePdf)
	rtr.HandleFunc(genUrl("{matrikel}/post/{postnr}"), handlers.HandlePosts)
	rtr.HandleFunc(genUrl("task/{tasknr}"), handlers.HandleTasks)
	http.Handle(genUrl(""), rtr)

	fmt.Printf("Start server at port%s:\n"+
		"   Coursname: %s\n"+
		"   Number of studentgroups: %v\n"+
		"   Course open: %v\n",
		conf.Port, conf.Course_name, conf.Group_number, conf.Open_course)

	//Start Server or exit with error message alter 5 seconds
	//err := http.ListenAndServeTLS(conf.Port, "../.tls/fullchain.pem", "../.tls/privkey.pem", nil)   template.Execute(writer, student1)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second * 5)
		os.Exit(1)
	}
}
