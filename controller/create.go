package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Vlad1slavZhuk/httpCRUD/data"
)

/*var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		panic("Set port!")
	}
}*/

// FormAdd - Load page add.html
func FormAdd(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v\n", r.RemoteAddr, r.Method, r.URL)
	http.ServeFile(w, r, "template/add.html")
}

// CreateCar - Adds a new ad to the database
func CreateCar(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v\n", r.RemoteAddr, r.Method, r.URL)
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "http://localhost:8081/", http.StatusMovedPermanently)
	}

	var car data.Car
	if r.Method == http.MethodPost {
		if r.Header.Get("Content-type") == "application/x-www-form-urlencoded" {
			brand := r.FormValue("brand")
			m := r.FormValue("model")
			color := r.FormValue("color")
			price, err := strconv.ParseFloat(r.FormValue("price"), 32)
			if err != nil {
				log.Printf("[ERROR] %v %v %v\n", r.RemoteAddr, r.Method, r.URL)
				http.Error(w, "Incorrect numbers.", http.StatusBadRequest)
				return
			}
			car = data.Car{
				Brand: brand,
				Model: m,
				Color: color,
				Price: float32(price),
			}
			if ok := data.AddCar(&car); !ok {
				log.Printf("[ERROR] %v %v %v\n", r.RemoteAddr, r.Method, r.URL)
				http.Error(w, "Error adding an ad for a car sale.", http.StatusBadRequest)
				return
			}
			fmt.Fprint(w, "(POST) SUCCESS! Added new car sale announcement.")
		} else {
			if err := car.FromJSON(r.Body); err != nil {
				log.Printf("[ERROR] %v %v %v\n", r.RemoteAddr, r.Method, r.URL)
				http.Error(w, "Error retrieving data from JSON.", http.StatusBadRequest)
				return
			}

			if ok := data.AddCar(&car); !ok {
				log.Printf("[ERROR] %v %v %v\n", r.RemoteAddr, r.Method, r.URL)
				http.Error(w, "Error adding an ad for a car sale.", http.StatusBadRequest)
				return
			}

			fmt.Fprint(w, "(JSON) SUCCESS! Added new car sale announcement.")
		}
	}
}
