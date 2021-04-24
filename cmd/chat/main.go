package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
		<html>
			<head>
				<title>Chat</title>
			</head>
			<body>
				Let's chat!
			</body>
		`))
	})

	listenAddr := "127.0.0.1:8080"
	log.Println("Starting server on ", listenAddr)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatal("Error starting web server:", err)
	}

}
