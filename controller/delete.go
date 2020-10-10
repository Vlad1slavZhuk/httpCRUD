package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Vlad1slavZhuk/httpCRUD/data"
	"github.com/gorilla/mux"
)

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v\n", r.RemoteAddr, r.Method, r.URL)
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Error in ParseUint.", http.StatusInternalServerError)
	}

	if ok := data.DeleteCar(uint(id)); !ok {
		http.Error(w, "Not found or was deleted.", http.StatusNotFound)
	} else {
		fmt.Fprintf(w, "(JSON) SUCCESS! Deleted ID = %v", id)
	}
}
