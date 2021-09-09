package server

import (
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	//fmt.Printf("listening at %s\n", port)
	return s.httpServer.ListenAndServe()
}
