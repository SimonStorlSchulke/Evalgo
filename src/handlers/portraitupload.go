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
		r.Body = http.MaxBytesReader(w, r.Body, 1*512*1024) // 500kb

		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Fprintf(w, "Fehler beim Upload des Portraits: %v Die Datei darf maximal 500kb groß sein.", err)
			return
		}
		ext := filepath.Ext(handler.Filename)
		if ext != ".jpg" && ext != ".JPG" && ext != ".jpeg" && ext != ".JPEG" {
			fmt.Fprint(w, "Fehler. Bild muss im .jpg Format sein")
			return
		}
		defer file.Close()

		f, err := os.OpenFile("./coursedata/portraits/"+fmt.Sprintf("%v.jpg", mat), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		st := fmt.Sprintf("<html><body>Neues Portrait hochgeladen <br>"+
			"<img src='./portraits/%v.jpg'><br>"+
			"<a href='./'>Zurück zur Übersicht </a></body></html>", mat)
		fmt.Fprintf(w, st)
		fmt.Printf("User %v updated his Portrait", mat)

		defer f.Close()
		io.Copy(f, file)
		http.Redirect(w, r, "./", http.StatusSeeOther)
	}
}
