package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	tpl      *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.tpl = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
		t.tpl.Execute(w, nil)
	})
}

func main() {
	http.Handle("/", &templateHandler{filename: "chat.html"})

	listenAddr := "127.0.0.1:8080"
	log.Println("Starting server on ", listenAddr)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatal("Error starting web server:", err)
	}

}
