package main

import (
	"log"
	"net/http"
	"os"
)

type server struct {
	mux *http.ServeMux
}

func NewServer() *server {
	s := &server{
		mux: http.NewServeMux(),
	}

	s.routes()
	return s
}

func main() {
	log.SetOutput(os.Stdout)
	log.Println("Starting Garage Door server...")

	s := NewServer()
	log.Fatal(http.ListenAndServe(":8090", s.mux))
}
