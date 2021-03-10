package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/ThomasThat467/INFO441Project/tree/main/servers/models/schedules"
)

// Schedule Handler ...
func (ctx *HandlerContext) ScheduleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			http.Error(w, "Plant request body must be json but got: %d", http.StatusUnsupportedMediaType)
			return
		} else {
			responseBody, _ := ioutil.ReadAll(r.Body)
			newSchedule := schedules.NewSchedule{}
			err := json.Unmarshal([]byte(responseBody), &newSchedule)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			//create new schedule
			createNewSchedule, err := newSchedule.ToSchedule()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			newScheduleInserted, err := ctx.ScheduleStore.Insert(createNewSchedule)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			addedSchedule, _ := json.Marshal(newScheduleInserted)
			w.Write(addedSchedule)
			return
		}
	} else {
		http.Error(w, "Method not allowed %d", http.StatusMethodNotAllowed)
		return
	}
}

// SpecificScheduleHandler ...
func (ctx *HandlerContext) SpecificScheduleHandler(w http.ResponseWriter, r *http.Request) {
	scheduleID := strings.TrimPrefix(r.URL.Path, "/v1/schedule/")
	currentSchedule := &schedules.Schedule{}

	if r.Method == http.MethodGet || r.Method == http.MethodPatch {
		scheduleID = strconv.FormatInt(currentSchedule.ID, 10)
		scheduleintID := currentSchedule.ID
		if r.Method == http.MethodGet {
			schedule, err := ctx.ScheduleStore.GetByID(scheduleintID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			} else {
				json, _ := json.Marshal(schedule)
				w.Header().Set("Content-Type", "application/json")
				w.Write(json)
				w.WriteHeader(http.StatusOK)
				return
			}
		} else if r.Method == http.MethodPatch {
			if strconv.FormatInt(currentSchedule.ID, 10) != scheduleID {
				fmt.Printf("Status not found for that id. Code: %d", http.StatusNotFound)
				return
			}
			if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
				fmt.Printf("Unnaccepted content type. Response body must be in JSON. Code: %d", http.StatusUnsupportedMediaType)
				return
			}
			marshaled, err := ioutil.ReadAll(r.Body)
			var updatesSchedInfo schedules.Updates
			if err == nil {
				json.Unmarshal([]byte(marshaled), &updatesSchedInfo)
			}
			updatedSchedule, err := ctx.ScheduleStore.Update(scheduleintID, &updatesSchedInfo)
			w.Header().Set("Content-Type", "application/json")
			marshalSchedule, err := json.Marshal(updatedSchedule)
			if err == nil {
				w.Write(marshalSchedule)
			}
			w.WriteHeader(http.StatusOK)
			return

		}
	}
}
