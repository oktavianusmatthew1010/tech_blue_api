package routes

import (
	"github.com/oktavianusmatthew1010/tech_blue_api/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/schedules", handlers.GetAllSchedules).Methods("GET")
	router.HandleFunc("/schedules/today", handlers.GetTodaySchedules).Methods("GET")
	router.HandleFunc("/schedules/{id}", handlers.GetScheduleDetails).Methods("GET")
	router.HandleFunc("/schedules/{id}/start", handlers.StartVisit).Methods("POST")
	router.HandleFunc("/schedules/{id}/end", handlers.EndVisit).Methods("POST")

	router.HandleFunc("/tasks/{taskId}/update", handlers.UpdateTask).Methods("POST")
}
