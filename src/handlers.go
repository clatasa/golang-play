package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

/*
	Here we manage stuff from the request and (if it's cool) pass it to the Repository
*/

var gopherManager = &dbGopherManager{}

func GetGophers(w http.ResponseWriter, r *http.Request) {

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = DB_LIMIT
	}

	gopherName := r.URL.Query().Get("gopher_name")

	gophersFound := gopherManager.findGophers(offset, limit, gopherName)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(gophersFound); err != nil {
		panic(err)
	}
}
func MakeGophers(w http.ResponseWriter, r *http.Request) {

	var gopher Gopher

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.Unmarshal(body, &gopher); err != nil {

		w.WriteHeader(http.StatusUnprocessableEntity) // unprocessable entity

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	} else {

		gopher.Id = int(gopherManager.createGopher(gopher))

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(gopher); err != nil {
			panic(err)
		}
	}

}
