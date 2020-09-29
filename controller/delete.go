package controller

import (
	"fmt"
	"github.com/Vlad1slavZhuk/httpCRUD/data"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

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