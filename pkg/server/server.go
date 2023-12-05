package server

import (
	"log"
	"net"
	"net/http"
	"path/filepath"
)

type Server struct {
	httpServer *http.Server
	doneChan   chan struct{}
}

func NewServer() *Server {
	currentDir, _ := filepath.Abs(filepath.Dir("."))
	staticPath := filepath.Join(currentDir, "static")

	fs := http.FileServer(http.Dir(staticPath))

	mux := http.NewServeMux()

	mux.Handle("/", fs)

	httpServer := &http.Server{
		Handler: mux,
	}

	return &Server{
		httpServer: httpServer,
		doneChan:   make(chan struct{}),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", ":3000")

	if err != nil {
		return err
	}

	go s.httpServer.Serve(ln)

	log.Print("Listening on :3000...")

	<-s.doneChan

	return nil
}

func (s *Server) Stop() error {
	s.httpServer.Close()

	close(s.doneChan)

	return nil
}
