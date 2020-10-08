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

func GetListCars(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v\n", r.RemoteAddr, r.Method, r.URL)
	w.Header().Set("Content-Type", "application/json")
	if len(data.GetListCars()) == 0 {
		http.Error(w, "List of cars is empty.", http.StatusBadRequest)
	} else {
		list, err := json.MarshalIndent(data.GetListCars(), "", "   ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, string(list))
	}
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v\n", r.RemoteAddr, r.Method, r.URL)
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "ID parsing error.", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if car, ok := data.GetCar(id); ok {
		c, err := json.MarshalIndent(car, "", "   ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, string(c))
	} else {
		http.Error(w, "Not Found.", http.StatusBadRequest)
	}
}
