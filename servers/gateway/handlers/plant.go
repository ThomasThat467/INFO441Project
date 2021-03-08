package handlers

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type PlantCotext struct {
}

func (ctx *PlantCotext) PlantHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			http.Error(w, "Platn request body must be json but got: %d", http.StatusUnsupportedMediaType)
			return
		} else {
			responseBody, _ := ioutil.ReadAll(r.Body)
			newPlant := plants.NewPlant{}

		}
	}
}

func (ctx *PlantCotext) SpecificPlantHandler(w http.ResponseWriter, r *http.Request) {
	plantID := userID := strings.TrimPrefix(r.URL.Path, "/v1/plant/")
	currentPlant := 

	if r.Method == http.MethodGet || r.Method == http.MethodPath {
		
	}
}