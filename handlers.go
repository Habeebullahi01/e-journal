package main

import (
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Checks /data for all the files it contains
// Reads the data within
// Returns a slice containing each as *Entry
func loadAllEntries(reqPath string) ([]Entry, error) {

	var e []Entry
	filepath.WalkDir("data", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			// Open the file and read the Entrydate
			// file, err := os.Open(path)
			// if err != nil {
			// 	return err
			// }
			// defer file.Close()
			var content *Entry
			b, _ := os.ReadFile(path)
			if json.Valid(b) {
				json.Unmarshal(b, &content)
				e = append(e, *content)
			}
		}
		return nil
	})
	return e, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the existence of /data
	if _, err := os.ReadDir("data"); err != nil {
		os.Mkdir("data", fs.FileMode(fs.ModeDir))
	}

	//Get names of all available entry files

	var availableEntries []Entry
	availableEntries, _ = loadAllEntries("")
	err := templates.ExecuteTemplate(w, "home.html", availableEntries)
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, path []string) {
	// e:= &EntryDate{Year: path[1], Month: path[2],}
	entry, err := load(path)
	if err != nil {
		log.Fatal(err)
		return
	}

	// fmt.Fprint(w, entry)
	renderTemplate(w, "view", entry)
	// e, nil:=load(path)
}

// func editHandler(w http.ResponseWriter, r *http.Request, title string) {
// 	fmt.Fprint(w, "editing "+title)
// }

// func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
// 	// create Entry object using the request body and call save() method
// 	// redirect to /view/entrydate if error is nil
// 	// redirect to /error/pastdate or /error/futuredate as necessitated by error
// }

func makeHandler(fn func(w http.ResponseWriter, r *http.Request, path []string)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		d := r.URL.Path
		ds := strings.Split(d, "/")
		ds = ds[2:]
		// if validPath.MatchString(strings.Join(ds, "/")) {

		// 	fn(w, r, ds)
		// } else {
		// 	http.Error(w, "Path malformed", http.StatusNotFound)
		// }
		// fmt.Println(ds)
		// m := validPath.FindStringSubmatch(r.URL.Path)
		// if m == nil {
		// 	http.NotFound(w, r)
		// 	// log.Fatal("m is empty!")
		// 	return
		fn(w, r, ds)
		// }
	}
}
