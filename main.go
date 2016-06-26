package main

import (
	"html/template"
	"log"
	"net/http"
	// "strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title       string
	Description string
	CreatedOn   time.Time
}

var templates map[string]*template.Template

func init() {

	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	// Compile templates
	templates["index"] = template.Must(template.ParseFiles("templates/index.html", "templates/layout.html"))
	templates["add"] = template.Must(template.ParseFiles("templates/add.html", "templates/layout.html"))
	templates["edit"] = template.Must(template.ParseFiles("templates/edit.html", "templates/layout.html"))
}

// render tmpls by name and add data
func renderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) {
	// ensure the template exists in the map
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The Template does not exits.", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var noteStore = make(map[string]Note)

var id int = 0

func getNotes(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", "layout", noteStore)
}
func addNote(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "add", "layout", noteStore)
}
func saveNote(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "save", "layout", noteStore)
}
func editNote(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "edit", "layout", noteStore)
}
func updateNote(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "update", "layout", noteStore)
}
func deleteNote(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "delete", "layout", noteStore)
}

// main
func main() {
	r := mux.NewRouter().StrictSlash(false)
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/", fs)
	r.HandleFunc("/", getNotes)
	r.HandleFunc("/notes/add", addNote)
	r.HandleFunc("/notes/save", saveNote)
	r.HandleFunc("/notes/edit/{id}", editNote)
	r.HandleFunc("/notes/update/{id}", updateNote)
	r.HandleFunc("/notes/delete/{id}", deleteNote)

	server := &http.Server{
		Addr:    ":9090",
		Handler: r,
	}

	log.Println("We're up on port 9090...")
	server.ListenAndServe()
}
