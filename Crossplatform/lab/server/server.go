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
	log.Print("Starting server on port " + s.port)
	log.Println("Endpoints: /solve")
	err := http.ListenAndServe(":"+s.port, nil)
	if err != nil {
		log.Println("Server closed: ", err.Error())
	}
	return nil
}
