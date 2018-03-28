package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Page struct für Einzelne Seite
type Page struct {
	Title string
	Body  []byte
}

func main() {
	//erzeuge Page struct
	p1 := &Page{"TestPage", []byte("This is a test Page")}
	//speicher als .txt
	p1.save()
	//lade Page von .txt
	p2, _ := loadPage("TestPage")
	//Print Page
	fmt.Println(string(p2.Body))
}

//speichert Page als .txt
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//lädt .txt als Page struct
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	//Print aktuellen URL Pfad. [1:] entfernt erstes Zeichen (hier: / )
	fmt.Fprintf(w, "Aktueller Pfad: %s", r.URL.Path[1:])
}
