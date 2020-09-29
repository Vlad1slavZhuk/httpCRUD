package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Vlad1slavZhuk/httpCRUD/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetListCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(data.GetListCars()) == 0 {
		http.Error(w, "List of cars is empty.", http.StatusBadRequest)
	} else {
		list, err := json.MarshalIndent(data.GetListCars(), "", "   ")
		if err != nil {
			log.Fatal(err)
		}
		w.Write(list)
	}
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	w.Header().Set("Content-Type", "application/json")
	if car, ok := data.GetCar(uint64(id)); ok {
		c, err := json.MarshalIndent(car, "", "   ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, string(c))
	} else {
		fmt.Fprint(w, "Not found.")
	}
}
