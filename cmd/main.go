package main

import (
	"context"
	handler "github.com/Vlad1slavZhuk/httpCRUD/controller"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.Hello)
	r.HandleFunc("/car", handler.FormAdd)
	r.HandleFunc("/car/add", handler.CreateCar).Methods(http.MethodPost, http.MethodGet)
	r.HandleFunc("/cars", handler.GetListCars).Methods(http.MethodGet)
	r.HandleFunc("/cars/{id:[0-9]+}", handler.GetCar).Methods(http.MethodGet)
	r.HandleFunc("/cars/{id:[0-9]+}", handler.UpdateCar).Methods(http.MethodPut, http.MethodPatch)
	r.HandleFunc("/cars/{id:[0-9]+}", handler.DeleteCar).Methods(http.MethodDelete)

	//log.Fatal(http.ListenAndServe(":8081", r))
	// create a new server
	s := http.Server{
		Addr:         ":8081",           // configure the bind address
		Handler:      r,                 // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
	}

	// start server
	go func() {
		log.Println("Starting server on port 8081")

		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
