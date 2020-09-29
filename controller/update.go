package controller

import (
	"fmt"
	"github.com/Vlad1slavZhuk/httpCRUD/data"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	var car data.Car
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Wrong data.", http.StatusBadRequest)
		return
	}

	if err := car.FromJSON(r.Body); err != nil {
		http.Error(w, "Wrong data.", http.StatusBadRequest)
		return
	}

	data.UpdateCar(id, &car)
	fmt.Fprintf(w, "Update a car!")
}
