package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	if descriptor := r.FormValue("descriptor"); descriptor != "" {
		log.Printf("Descriptor: %s", r.FormValue("descriptor"))
	}
	if animal1 := r.FormValue("animal1"); animal1 != "" {
		log.Printf("Animal1: %s", r.FormValue("animal1"))
	}
	if verb := r.FormValue("verb"); verb != "" {
		log.Printf("Verb: %s", r.FormValue("verb"))
	}
	if adverb := r.FormValue("adverb"); adverb != "" {
		log.Printf("Adverb: %s", r.FormValue("adjective"))
	}
	if adjective := r.FormValue("adjective"); adjective != "" {
		log.Printf("Adjective: %s", r.FormValue("adjective"))
	}
	if noun := r.FormValue("noun"); noun != "" {
		log.Printf("noun: %s", r.FormValue("noun"))
	}
	err := t.templ.Execute(w, r)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.Handle("/", &templateHandler{filename: "index.html"})
	http.Handle("/madlib", &templateHandler{filename: "templates/madlib.html"})
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
