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
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
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
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	err := t.tpl.Execute(w, data)
	if err != nil {
		log.Println("ERROR: " + err.Error())
	}
}

func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "The address application listens on")
	tracingEnabled := flag.Bool("trace", false, "Enable tracing")
	flag.Parse()

	gomniauth.SetSecurityKey("some secret string here")
	gomniauth.WithProviders(
		github.New("abc3d1f9df5e9712ef6b", os.Getenv("GITHUB_OAUTH2_SECRET"), "http://"+*addr+"/auth/callback/github"),
		google.New("google client", "google secret", "http://"+*addr+"/auth/callback/goole"),
	)

	r := NewRoom()
	if *tracingEnabled {
		r.tracer = trace.New(os.Stdout)
	}

	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)

	http.Handle("/assets/",
		http.StripPrefix("/assets",
			http.FileServer((http.Dir("./assets"))),
		))
	http.Handle("/room", r)
	go r.run()

	log.Println("Starting server on ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Error starting web server:", err)
	}

}
