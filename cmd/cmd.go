package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"weproov/internal/images"

	"github.com/gorilla/mux"
)

type Server struct {
	Addr string
	Router *mux.Router

}

func NewServer(addr string) *Server {
	return &Server{
		Addr: addr,
	}
}

func (s *Server) InitServer() *http.Server {

	return &http.Server{
		Addr:          	s.Addr,
		Handler:        s.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
}

func (s *Server) Execute() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	s.Router = mux.NewRouter()

	imgServices := images.NewImagesServices()
	imgRoutes := images.NewImagesRoutes(s.Router, imgServices)
	imgRoutes.Routes()

	srv := s.InitServer()
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	go func() {
		<-quit
		logger.Println("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	log.Println("Server is runnning")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", ":8083", err)
	}
	<-done
	logger.Println("Server stopped")
}