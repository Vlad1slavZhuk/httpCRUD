package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Vlad1slavZhuk/httpCRUD/model"
	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func FormAdd(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "template/add.html")
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	var car model.Car
	if r.Method == http.MethodGet {
		err := json.NewDecoder(r.Body).Decode(&car)
		if err != nil {
			fmt.Fprintln(w, err)
		}
		fmt.Fprint(w, car)
		model.AddCar(&car)
	}
	if r.Method == http.MethodPost {
		m := r.FormValue("model")
		color := r.FormValue("color")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		if m == "" || color == "" || price == 0 {
			fmt.Fprint(w, "Error.")
			return
		}
		car = model.Car{
			Model: m,
			Color: color,
			Price: price,
		}
		model.AddCar(&car)
		fmt.Fprint(w, "Add a new car.")
	}
	//TODO add decode json
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	ok := model.DeleteCar(id)
	if !ok {
		fmt.Fprint(w, "Error.")
	} else {
		fmt.Fprint(w, "Delete a car!")
	}
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update a car!")
}

func GetListCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(model.GetListCars()) == 0 {
		fmt.Fprintf(w, "null")
	} else {
		json, err := json.MarshalIndent(model.GetListCars(), "", "   ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, string(json))
	}
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	w.Header().Set("Content-Type", "application/json")
	if car, ok := model.GetCar(uint64(id)); ok {
		c, err := json.MarshalIndent(car, "", "   ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, string(c))
	} else {
		fmt.Fprint(w, "Not found.")
	}
}
