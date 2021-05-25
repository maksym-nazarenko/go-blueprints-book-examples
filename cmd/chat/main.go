package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

const (
	TEMPLATES_DIR_ENV_NAME = "TEMPLATES_DIR"
)

type templateHandler struct {
	once     sync.Once
	filename string
	tpl      *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		var templatesDir string
		var ok bool
		if templatesDir, ok = os.LookupEnv(TEMPLATES_DIR_ENV_NAME); !ok {
			templatesDir = "templates"
		}

		t.tpl = template.Must(template.ParseFiles(filepath.Join(templatesDir, t.filename)))
	})
	t.tpl.Execute(w, nil)
}

func main() {
	r := NewRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	go r.run()

	listenAddr := "127.0.0.1:8080"
	log.Println("Starting server on ", listenAddr)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatal("Error starting web server:", err)
	}

}
