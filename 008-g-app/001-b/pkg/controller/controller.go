package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

)


func GetAllData(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := mh.service.GetMovies()
	if err != nil {
		http.Error(w, "Unable to get all data", http.StatusInternalServerError)
		return
	}

	db, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(db)
}


func GetData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))

	data, err := mh.service.GetData(id)
	if err != nil {
		if errors.Is(err, service.ErrIDIsNotValid) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if errors.Is(err, service.ErrMovieNotFound) { 
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(db)
}

func CreateData(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Movie is successfully created"))
}


func UpdateData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))


	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Successfully Updated"))
}

func DeleteData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("data has been deleted succesfully"))
}

func DeleteAllData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("all data successfully deleted"))
}
