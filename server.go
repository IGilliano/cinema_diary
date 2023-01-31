package cinema_diary

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpserver *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpserver = &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		MaxHeaderBytes:    1 << 20,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	return s.httpserver.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpserver.Shutdown(ctx)
}
