package main

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("/(edit|view|save)/([a-zA-Z0-9]+)$")

// var validViewPath = regexp.MustCompile("/view/([0-9])")

var templates = template.Must(template.ParseFiles("templates/view.html", "templates/home.html"))

// func load(entryDate EntryDate) (*Entry, error) {
// 	// "title" can be a map of day,month and year
// 	// os.Open("/data/"+title.year+"/"+title.month+"/"+title.day+".json")
// 	// Read the json file and create Entry instance using the key-value pairs
// 	file, err := os.Open("/data/" + string(entryDate.Year) + "/" + entryDate.Month + "/" + string(entryDate.Day) + ".txt")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()
// 	body, err := os.ReadFile(file.Name())
// 	if err != nil {
// 		return nil, err
// 	}
// 	// var date EntryDate
// 	// date.FullDate = time.Now()
// 	// date.Year, _, date.Day = date.FullDate.EntryDate()
// 	// date.Month = date.FullDate.Month().String()
// 	// date.Weekday = date.FullDate.Weekday().String()
// 	title := string(entryDate.Day) + "-" + entryDate.Weekday
// 	e := &Entry{Title: title, Date: entryDate, Body: string(body)}
// 	return e, nil
// }

// func renderTemplate(w http.ResponseWriter, name string, data *Entry) {

// 	templates.ExecuteTemplate(w, name+".html", data)
// }

func main() {
	// File Organization in Golang
	http.HandleFunc("/", homeHandler)
	http.Handle("/view/", makeHandler(viewHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
