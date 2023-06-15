package main

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

const addr = ":8080" // TODO Make this configurable.

var templates *template.Template
var me = Person{
	Name:    "Bas",
	Hobbies: []string{"Gaming", "Reading", "Extreme Ironing"},
}

type Person struct {
	Name    string
	Hobbies []string
}

func init() {
	templatesPath := filepath.Join("templates", "*.html")
	templates = template.Must(template.New("").ParseGlob(templatesPath))
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/about/", handleAbout)
	http.ListenAndServe(addr, nil)

}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		log.Println("error:", err)
	}
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "about.html", me); err != nil {
		log.Println("error:", err)
	}
}
