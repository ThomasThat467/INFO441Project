package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/ThomasThat467/INFO441Project/tree/main/servers/models/plants"
)

//not sure we need plant context
type PlantCotext struct {
	PlantStore plants.Store
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
			err := json.Unmarshal([]byte(responseBody), &newPlant)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			//create new plant
			createNewPlant, err := newPlant.ToPlant()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			newPlantInserted, err := ctx.PlantStore.Insert(createNewPlant)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			addedPlant, _ := json.Marshal(newPlantInserted)
			w.Write(addedPlant)
			return
		}
	} else {
		http.Error(w, "Method not allowed %d", http.StatusMethodNotAllowed)
		return
	}
}

func (ctx *PlantCotext) SpecificPlantHandler(w http.ResponseWriter, r *http.Request) {
	plantID := strings.TrimPrefix(r.URL.Path, "/v1/plant/")
	currentPlant := &plants.Plant{}

	if r.Method == http.MethodGet || r.Method == http.MethodPatch {
		plantID = strconv.FormatInt(currentPlant.ID, 10)
		plantintID := currentPlant.ID
		if r.Method == http.MethodGet {
			plant, err := ctx.PlantStore.GetByID(plantintID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			} else {
				json, _ := json.Marshal(plant)
				w.Header().Set("Content-Type", "application/json")
				w.Write(json)
				w.WriteHeader(http.StatusOK)
				return
			}
		} else if r.Method == http.MethodPatch {
			if strconv.FormatInt(currentPlant.ID, 10) != plantID {
				fmt.Printf("Status not found for that id. Code: %d", http.StatusNotFound)
				return
			}
			if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
				fmt.Printf("Unnaccepted content type. Response body must be in JSON. Code: %d", http.StatusUnsupportedMediaType)
				return
			}
			marshaled, err := ioutil.ReadAll(r.Body)
			var updatesPlantInfo plants.Updates
			if err == nil {
				json.Unmarshal([]byte(marshaled), &updatesPlantInfo)
			}
			updatedPlant, err := ctx.PlantStore.Update(plantintID, &updatesPlantInfo)
			w.Header().Set("Content-Type", "application/json")
			marshalPlant, err := json.Marshal(updatedPlant)
			if err == nil {
				w.Write(marshalPlant)
			}
			w.WriteHeader(http.StatusOK)
			return

		}
	}
	//not sure delte part
	if r.Method == http.MethodDelete {
		plantID = strconv.FormatInt(currentPlant.ID, 10)
		plantintID := currentPlant.ID
		if strconv.FormatInt(currentPlant.ID, 10) != plantID {
			fmt.Printf("Status not found for that id. Code: %d", http.StatusNotFound)
			return
		}
		if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
			fmt.Printf("Unnaccepted content type. Response body must be in JSON. Code: %d", http.StatusUnsupportedMediaType)
			return
		}
		deletePlant, err := ctx.PlantStore.Delete(plantintID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else {
			json, _ := json.Marshal(deletePlant)
			w.Write(json)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
}
