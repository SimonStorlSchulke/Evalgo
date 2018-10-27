package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"./courseconfig"
	"./handlers"
	"github.com/gorilla/mux"
)

var conf courseconfig.Config = courseconfig.GetConfig()

func genUrl(str string) string {
	return fmt.Sprintf("%s%s", conf.Root_url, str)
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {

	rtr := mux.NewRouter()

	FileServer := http.FileServer(http.Dir("static"))
	http.Handle(genUrl("static/"), http.StripPrefix(genUrl("static/"), FileServer))

	FileServerRes := http.FileServer(http.Dir("coursedata/ressources"))
	http.Handle(genUrl("res/"), http.StripPrefix(genUrl("res/"), FileServerRes))

	FileServerPortraits := http.FileServer(http.Dir("coursedata/portraits"))
	http.Handle(genUrl("portraits/"), http.StripPrefix(genUrl("portraits/"), FileServerPortraits))

	FileServerImg := http.FileServer(http.Dir("coursedata/tasks/img"))
	http.Handle(genUrl("img/"), http.StripPrefix(genUrl("img/"), FileServerImg))

	rtr.HandleFunc(genUrl("register"), handlers.HandleRegister)
	rtr.HandleFunc(genUrl("login"), handlers.HandleLogin)
	rtr.HandleFunc(genUrl("authlogin"), handlers.HandleAuthLogin)
	rtr.HandleFunc(genUrl("uploadportrait"), handlers.PortraitUpload)
	rtr.HandleFunc(genUrl("post"), handlers.HandlePostForm)
	rtr.HandleFunc(genUrl("process"), handlers.ProcessPost)
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

	fmt.Printf("Course: %s\n"+
		"Course open: %v\n"+
		"Groups: %v\n"+
		"Cards: %v\n"+
		"Grades: %v \n",
		conf.Course_name, conf.Open_course, conf.Group_number, conf.Enable_cards, conf.Enable_grades)

	var localhost bool
	if len(os.Args) > 1 {
		if os.Args[1] == "localhost" {
			localhost = true
		}
	}

	if localhost {
		fmt.Printf("Started server at localhost%s%s\n",
			conf.Port, conf.Root_url)

		//Start Server or exit with error message after 5 seconds
		err := http.ListenAndServe(":8080", Log(http.DefaultServeMux))
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second * 5)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Started server at /%s%s\n",
			conf.Port, conf.Root_url)

		//Start Server or exit with error message after 5 seconds
		err := http.ListenAndServeTLS(conf.Port, "../.tls/fullchain.pem", "../.tls/privkey.pem", nil)
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second * 5)
			os.Exit(1)
		}
	}
}
