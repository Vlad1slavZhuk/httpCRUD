package main

import (
	"log"
	"net/http"

	handler "github.com/Vlad1slavZhuk/httpCRUD/controller"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.Hello)
	r.HandleFunc("/cars", handler.GetListCars).Methods(http.MethodGet)
	r.HandleFunc("/cars/{id}", handler.GetCar).Methods(http.MethodGet)
	r.HandleFunc("/car", handler.FormAdd)
	r.HandleFunc("/car/add", handler.CreateCar).Methods(http.MethodPost, http.MethodGet)
	r.HandleFunc("/cars/{id}", handler.UpdateCar).Methods(http.MethodPut)
	r.HandleFunc("/cars/{id}", handler.DeleteCar).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8081", r))
}
