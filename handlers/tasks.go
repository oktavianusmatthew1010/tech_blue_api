package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID := params["taskId"]

	var req struct {
		Completed bool   `json:"completed"`
		Reason    string `json:"reason,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	for i, s := range schedules {
		for j, t := range s.Tasks {
			if t.ID == taskID {
				schedules[i].Tasks[j].Completed = req.Completed
				if !req.Completed && req.Reason != "" {
					schedules[i].Tasks[j].Reason = req.Reason
				} else {
					schedules[i].Tasks[j].Reason = ""
				}
				respondWithJSON(w, http.StatusOK, schedules[i])
				return
			}
		}
	}

	respondWithError(w, http.StatusNotFound, "Task not found")
}
