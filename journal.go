package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// var validPath = regexp.MustCompile("/(edit|view|save)/([0-9]+)/([A-Z][a-zA-Z])/(2[0-9]-[A-Z][a-zA-Z])$")

// var validViewPath = regexp.MustCompile("/view/([0-9])")

var templates = template.Must(template.ParseFiles("templates/view.html", "templates/home.html"))

func load(path []string) (*Entry, error) {
	// "title" can be a map of day,month and year
	// os.Open("/data/"+title.year+"/"+title.month+"/"+title.day+".json")
	// Read the json file and create Entry instance using the key-value pairs
	// file, err := os.Open("data/" + strings.Join(path, "/") + ".json")
	// if err != nil {
	// 	return nil, err
	// }
	// defer file.Close()
	var content *Entry
	b, err := os.ReadFile("data/" + strings.Join(path, "/") + ".json")
	if err != nil {
		fmt.Println(err)
	}
	if json.Valid(b) {
		json.Unmarshal(b, &content)
	}
	return content, nil
}

func renderTemplate(w http.ResponseWriter, name string, data *Entry) {

	templates.ExecuteTemplate(w, name+".html", data)
}

func main() {
	// File Organization in Golang
	http.HandleFunc("/", homeHandler)
	http.Handle("/view/", makeHandler(viewHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
