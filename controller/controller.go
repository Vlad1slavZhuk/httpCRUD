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

func FormAdd(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-type", "application/x-www-form-urlencoded")
	http.ServeFile(w, r, "template/add.html")
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "http://localhost:8081/", http.StatusMovedPermanently)
	}

	var car data.Car
	if r.Method == http.MethodPost {
		if r.Header.Get("Content-type") == "application/x-www-form-urlencoded" {
			m := r.FormValue("model")
			color := r.FormValue("color")
			price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
			car = data.Car{
				Model: m,
				Color: color,
				Price: price,
			}
			if ok := data.AddCar(&car); !ok {
				http.Error(w, "Wrong data.", http.StatusBadRequest)
				return
			}
			fmt.Fprint(w, "Add a new car.")
		} else {
			if err := car.FromJSON(r.Body); err != nil {
				http.Error(w, "Wrong data.", http.StatusBadRequest)
				return
			}

			if ok := data.AddCar(&car); !ok {
				http.Error(w, "Wrong data.", http.StatusBadRequest)
				return
			}

			fmt.Fprint(w, "Add a new car.")
		}
	}
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Error in ParseUint.", http.StatusInternalServerError)
	}
	ok := data.DeleteCar(id)
	if !ok {
		http.Error(w, "Not found or already deleted.", http.StatusNotFound)
	} else {
		fmt.Fprint(w, "Delete a car!")
	}
}

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
