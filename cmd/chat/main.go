package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/maxim-nazarenko/go-blueprints-book-examples/trace"
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
	err := t.tpl.Execute(w, r)
	if err != nil {
		log.Println("ERROR: " + err.Error())
	}
}

func main() {
	r := NewRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	go r.run()

	addr := flag.String("addr", "127.0.0.1:8080", "The address application listens on")
	flag.Parse()

	log.Println("Starting server on ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Error starting web server:", err)
	}

}
