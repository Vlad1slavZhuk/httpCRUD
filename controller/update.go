package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Vlad1slavZhuk/httpCRUD/data"
	"github.com/gorilla/mux"
)

// UpdateCar...
func UpdateCar(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v\n", r.RemoteAddr, r.Method, r.URL)
	var car data.Car
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "ID parsing error.", http.StatusBadRequest)
		return
	}

	if err := car.FromJSON(r.Body); err != nil {
		http.Error(w, "Error retrieving data from JSON.", http.StatusBadRequest)
		return
	}

	if ok := data.UpdateCar(id, &car); !ok {
		http.Error(w, "Data update error.", http.StatusBadRequest)
	} else {
		fmt.Fprintf(w, "(JSON) ID %v UPDATED!", id)
	}

}
