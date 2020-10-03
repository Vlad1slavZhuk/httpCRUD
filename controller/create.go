package controller

import (
	"fmt"
	"github.com/Vlad1slavZhuk/httpCRUD/data"
	"net/http"
	"os"
	"strconv"
)

var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		 panic("Set port!")
	}
}

func FormAdd(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "template/add.html")
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "http://localhost:"+port+"/", http.StatusMovedPermanently)
	}

	var car data.Car
	if r.Method == http.MethodPost {
		if r.Header.Get("Content-type") == "application/x-www-form-urlencoded" {
			brand := r.FormValue("brand")
			m := r.FormValue("model")
			color := r.FormValue("color")
			price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
			car = data.Car{
				Brand: brand,
				Model: m,
				Color: color,
				Price: price,
			}
			if ok := data.AddCar(&car); !ok {
				http.Error(w, "Error adding a sell order.", http.StatusBadRequest)
				return
			}
			fmt.Fprint(w, "(POST) SUCCESS! Added new car sale announcement.")
		} else {
			if err := car.FromJSON(r.Body); err != nil {
				http.Error(w, "Error retrieving data from JSON.", http.StatusBadRequest)
				return
			}

			if ok := data.AddCar(&car); !ok {
				http.Error(w, "Error adding a sell order.", http.StatusBadRequest)
				return
			}

			fmt.Fprint(w, "(JSON) SUCCESS! Added new car sale announcement.")
		}
	}
}
