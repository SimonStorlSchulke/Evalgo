package handlers

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
	"time"
)

func PortraitUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("p method:", r.Method)
	isLoggedIn, mat := loggedIn(r)
	if !isLoggedIn {
		fmt.Fprintf(w, "loggin to change your portrait")
		return
	}

	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("./templates/portrait-upload.go.html")
		t.Execute(w, token)
	} else {
		fmt.Println("p method:", r.Method)
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		ext := filepath.Ext(handler.Filename)
		if ext != ".png" {
			fmt.Fprintf(w, "not an png file but %v", ext)
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "uploaded new Portrait")
		f, err := os.OpenFile("./coursedata/portraits/"+fmt.Sprintf("%v.png", mat), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()
		io.Copy(f, file)
		http.Redirect(w, r, "./", http.StatusSeeOther)
	}
}
