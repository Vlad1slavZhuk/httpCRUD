package data

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

// NewServer - create and return new server.
func NewServer(r *mux.Router, port string) *http.Server {
	return &http.Server{
		Addr:         port,             // configure the bind address
		Handler:      r,                // set the default handler
		ReadTimeout:  5 * time.Second,  // max time to read request from the client
		WriteTimeout: 10 * time.Second, // max time to write response to the client
	}
}

// Run - start server.
func Run(s *http.Server) {
	log.Println("Starting server on port", s.Addr)

	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}

	Shutdown(s)
}

// Shutdown - shutdown server after get signal - "Interrupt"
func Shutdown(s *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)
	defer close(c)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	if err := s.Shutdown(ctx); err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
	}
}
