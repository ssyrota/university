package main

import (
	"log"
	"net/http"
)

func NewServer(port string) Server {
	return Server{port: port}
}

type Server struct {
	port string
}

func (s Server) Listen() error {
	http.HandleFunc("/solve", SolvePost)
	http.HandleFunc("/det", DetPost)

	fs := http.FileServer(http.Dir("../web_client"))
	http.Handle("/", fs)

	log.Println(`
[INFO] Starting server on port ` + s.port + `
[INFO] Endpoints: /solve, /det
[INFO] Website: ./index.html
[INFO] Open in browser: http://localhost:` + s.port + `/index.html`,
	)

	err := http.ListenAndServe(":"+s.port, nil)
	if err != nil {
		log.Println("[ERROR] Server closed: ", err.Error())
	}
	return nil
}
