package handlers

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
