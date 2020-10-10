package main

import (
	"net/http"
	"os"

	handler "github.com/Vlad1slavZhuk/httpCRUD/controller"
	server "github.com/Vlad1slavZhuk/httpCRUD/data"
	"github.com/gorilla/mux"
)

var port string

func init() {
	port = os.Getenv("PORT")

	if port == "" {
		panic("Set port!")
	}
}

func main() {
	r := mux.NewRouter()

	//Load page add.html
	r.HandleFunc("/", handler.FormAdd)
	//GET
	r.HandleFunc("/cars", handler.GetListCars).Methods(http.MethodGet)
	r.HandleFunc("/cars/{id:[0-9]+}", handler.GetCar).Methods(http.MethodGet)
	//POST
	r.HandleFunc("/car", handler.CreateCar).Methods(http.MethodPost, http.MethodGet)
	//PUT
	r.HandleFunc("/cars/{id:[0-9]+}", handler.UpdateCar).Methods(http.MethodPut)
	//DELETE
	r.HandleFunc("/cars/{id:[0-9]+}", handler.DeleteCar).Methods(http.MethodDelete)

	// Create a new server
	s := server.NewServer(r, ":"+port)
	// Start server
	server.Run(s)
}
