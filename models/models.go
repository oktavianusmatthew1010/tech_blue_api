package models

import "time"

type Schedule struct {
	ID           string     `json:"id"`
	CaregiverID  string     `json:"caregiverId"`
	ClientName   string     `json:"clientName"`
	ServiceName  string     `json:"serviceName"`
	ImageUrl     string     `json:"imageUrl"`
	EmailAddress string     `json:"emailAddress"`
	Phone        string     `json:"phone"`
	DateService  time.Time  `json:"dateService"`
	StartTime    time.Time  `json:"startTime"`
	EndTime      time.Time  `json:"endTime"`
	ServiceNote  string     `json:"serviceNote"`
	Location     Location   `json:"location"`
	Status       string     `json:"status"`
	Tasks        []Task     `json:"tasks"`
	StartVisit   *VisitTime `json:"startVisit,omitempty"`
	EndVisit     *VisitTime `json:"endVisit,omitempty"`
}

type Task struct {
	ID          string `json:"id"`
	ScheduleID  string `json:"scheduleId"`
	TaskName    string `json:"taskName"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Reason      string `json:"reason,omitempty"`
}

type Location struct {
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type VisitTime struct {
	Timestamp time.Time `json:"timestamp"`
	Location  Location  `json:"location"`
}
