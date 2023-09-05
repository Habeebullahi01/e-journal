package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

type Entry struct {
	Title string
	Date  string
	Keys  []string
	Body  string
}

var validPath = regexp.MustCompile("/(edit|view|save)/([a-zA-Z0-9]+)$")
var validViewPath = regexp.MustCompile("/view/([0-9])")
var templates = template.Must(template.ParseFiles("templates/view.html"))

func (e *Entry) save() error {
	// check if present date corresponds to the entry date, only allow saving when true to prevent alteration of the file after the day has passed.
	// also check that entry date is not in the future before saving
	//save as a json file, with the day of the month(from the entry date) as the name, in the folders corresponding to the month and year (also gotten from the entry date). Create "title", "date", "keys" and "content" properties with their corresponding values.
	return nil
}

func load(title string) (*Entry, error) {
	// "title" can be a map of day,month and year
	// os.Open("/data/"+title.year+"/"+title.month+"/"+title.day+".json")
	// Read the json file and create Entry instance using the key-value pairs
	file, err := os.Open("/data/" + title + ".txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	body, _ := os.ReadFile(file.Name())
	e := &Entry{Title: title, Date: time.DateOnly, Body: string(body)}
	return e, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	entry, err := load(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	fmt.Fprint(w, entry)
	renderTemplate(w, "view", entry)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Fprint(w, "editing "+title)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// create Entry object using the request body and call save() method
	// redirect to /view/entrydate if error is nil
	// redirect to /error/pastdate or /error/futuredate as necessitated by error
}

func makeHandler(fn func(w http.ResponseWriter, r *http.Request, title string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			// log.Fatal("m is empty!")
			return
		}
		fn(w, r, m[2])
	}
}

func renderTemplate(w http.ResponseWriter, name string, data *Entry) {

	templates.ExecuteTemplate(w, "templates/"+name+".html", data)
}

func main() {
	// File Organization in Golang

	http.Handle("/view/", makeHandler(viewHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
