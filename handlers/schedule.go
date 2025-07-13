package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/oktavianusmatthew1010/tech_blue_api/models"

	"github.com/gorilla/mux"
)

var schedules = models.GenerateMockSchedules()

func GetAllSchedules(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, schedules)
}

func GetTodaySchedules(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	var todaySchedules []models.Schedule
	for _, s := range schedules {
		if s.StartTime.After(today) && s.StartTime.Before(today.Add(24*time.Hour)) {
			todaySchedules = append(todaySchedules, s)
		}
	}

	respondWithJSON(w, http.StatusOK, todaySchedules)
}

func GetScheduleDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	scheduleID := params["id"]

	for _, s := range schedules {
		if s.ID == scheduleID {
			respondWithJSON(w, http.StatusOK, s)
			return
		}
	}

	respondWithError(w, http.StatusNotFound, "Schedule not found")
}

func StartVisit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	scheduleID := params["id"]

	var req struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	for i, s := range schedules {
		if s.ID == scheduleID {
			schedules[i].Status = "in-progress"
			schedules[i].StartVisit = &models.VisitTime{
				Timestamp: time.Now(),
				Location: models.Location{
					Latitude:  req.Latitude,
					Longitude: req.Longitude,
				},
			}
			respondWithJSON(w, http.StatusOK, schedules[i])
			return
		}
	}

	respondWithError(w, http.StatusNotFound, "Schedule not found")
}

func EndVisit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	scheduleID := params["id"]

	var req struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	for i, s := range schedules {
		if s.ID == scheduleID {
			schedules[i].Status = "completed"
			schedules[i].EndVisit = &models.VisitTime{
				Timestamp: time.Now(),
				Location: models.Location{
					Latitude:  req.Latitude,
					Longitude: req.Longitude,
				},
			}
			respondWithJSON(w, http.StatusOK, schedules[i])
			return
		}
	}

	respondWithError(w, http.StatusNotFound, "Schedule not found")
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
