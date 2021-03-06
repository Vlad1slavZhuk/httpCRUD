package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Vlad1slavZhuk/httpCRUD/data"
	"github.com/gorilla/mux"
)

// GetListCars - Receives all ads for car sales.
func GetListCars(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v\n", r.RemoteAddr, r.Method, r.URL)
	w.Header().Set("Content-Type", "application/json")
	if len(data.GetListCars()) == 0 {
		log.Printf("[ERROR] %v %v %v\n", r.RemoteAddr, r.Method, r.URL)
		http.Error(w, "List of cars is empty.", http.StatusBadRequest)
	} else {
		list, err := json.MarshalIndent(data.GetListCars(), "", "   ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, string(list))
	}
}

// GetCar - Receives an ad for the sale of a car by ID
func GetCar(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v\n", r.RemoteAddr, r.Method, r.URL)
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		log.Printf("[ERROR] %v %v %v\n", r.RemoteAddr, r.Method, r.URL)
		http.Error(w, "ID parsing error.", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if car, ok := data.GetCar(uint(id)); ok {
		c, err := json.MarshalIndent(car, "", "   ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, string(c))
	} else {
		log.Printf("[ERROR] %v %v %v\n", r.RemoteAddr, r.Method, r.URL)
		http.Error(w, "Not Found.", http.StatusNotFound)
	}
}
