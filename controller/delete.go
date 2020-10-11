package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Vlad1slavZhuk/httpCRUD/data"
	"github.com/gorilla/mux"
)

// DeleteCar - Removes ad by ID. After deletion, sorts the ID of the rest of the ad.
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v\n", r.RemoteAddr, r.Method, r.URL)
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		log.Printf("[ERROR] %v %v %v\n", r.RemoteAddr, r.Method, r.URL)
		http.Error(w, "Error in ParseUint.", http.StatusInternalServerError)
	}

	if ok := data.DeleteCar(uint(id)); !ok {
		log.Printf("[ERROR] %v %v %v\n", r.RemoteAddr, r.Method, r.URL)
		http.Error(w, "Not found or was deleted.", http.StatusNotFound)
	} else {
		fmt.Fprintf(w, "(JSON) SUCCESS! Deleted ID = %v", id)
	}
}
