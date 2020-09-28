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

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\nHello World!")
}

func FormAdd(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "template/add.html")
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	var car data.Car
	if r.Method == http.MethodGet {

		//err := json.NewDecoder(r.Body).Decode(&car)
		if err := car.FromJSON(r.Body); err != nil {
			http.Error(w, "Wrong data.", http.StatusBadRequest)
			return
		}
		fmt.Fprint(w, car)
		data.AddCar(&car)
	}
	if r.Method == http.MethodPost {
		m := r.FormValue("model")
		color := r.FormValue("color")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		if m == "" || color == "" || price == 0 {
			http.Error(w, "Wrong data.", http.StatusBadRequest)
			return
		}
		car = data.Car{
			Model: m,
			Color: color,
			Price: price,
		}
		data.AddCar(&car)
		fmt.Fprint(w, "Add a new car.")
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
		http.Error(w, "Not found and already deleted.", http.StatusInternalServerError)
	} else {
		fmt.Fprint(w, "Delete a car!")
	}
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {

	//TODO
	var car data.Car
	if r.Method == http.MethodPut {
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
		fmt.Fprint(w, "\n"+string(c))
	} else {
		fmt.Fprint(w, "\nNot found.")
	}
}
