package models

import (
	"time"
)

func GenerateMockSchedules() []Schedule {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	return []Schedule{
		{
			ID:           "1",
			CaregiverID:  "caregiver1",
			ClientName:   "John Doe",
			ServiceName:  "Service 1",
			ImageUrl:     "https://static.vecteezy.com/system/resources/previews/024/183/525/non_2x/avatar-of-a-man-portrait-of-a-young-guy-illustration-of-male-character-in-modern-color-style-vector.jpg",
			EmailAddress: "john.doe@gmail.com",
			Phone:        "+628119211231",
			DateService:  today,
			StartTime:    today.Add(9 * time.Hour),
			EndTime:      today.Add(11 * time.Hour),
			ServiceNote:  "Client reports onset of anxiety at start of freshman year of college. Reports on and off mild anxiety prior to that. Reports intrusive thoughts of failing classes.",
			Location: Location{
				Address:   "123 Main St, Anytown, USA",
				Latitude:  34.052235,
				Longitude: -118.243683,
			},
			Status: "upcoming",
			Tasks: []Task{
				{ID: "1-1", ScheduleID: "1", TaskName: "Medicine", Description: "Give medication to client by daily base", Completed: false},
				{ID: "1-2", ScheduleID: "1", TaskName: "Cleaning", Description: "Assist Client with bathing, new cloth and cleaning the bed", Completed: false},
			},
		},
		{
			ID:          "2",
			CaregiverID: "caregiver1",
			ClientName:  "Jane Smith",
			ServiceName: "Service 2",
			ImageUrl:    "https://static.vecteezy.com/system/resources/previews/024/183/525/non_2x/avatar-of-a-man-portrait-of-a-young-guy-illustration-of-male-character-in-modern-color-style-vector.jpg",
			DateService: today,
			StartTime:   today.Add(13 * time.Hour),
			EndTime:     today.Add(15 * time.Hour),
			ServiceNote: "Client reports onset of anxiety at start of freshman year of college. Reports on and off mild anxiety prior to that. Reports intrusive thoughts of failing classes.",
			Location: Location{
				Address:   "456 Oak Ave, Somewhere, USA",
				Latitude:  34.052235,
				Longitude: -118.253683,
			},
			Status: "upcoming",
			Tasks: []Task{
				{ID: "2-1", ScheduleID: "2", TaskName: "Lunch", Description: "Prepare lunch for Client", Completed: false},
				{ID: "2-2", ScheduleID: "2", TaskName: "Cleaning", Description: "Light housekeeping", Completed: false},
			},
		},
		{
			ID:          "3",
			CaregiverID: "caregiver1",
			ClientName:  "Robert Johnson",
			ServiceName: "Service 3",
			ImageUrl:    "https://static.vecteezy.com/system/resources/previews/024/183/525/non_2x/avatar-of-a-man-portrait-of-a-young-guy-illustration-of-male-character-in-modern-color-style-vector.jpg",
			DateService: today,
			StartTime:   today.Add(-2 * time.Hour),
			EndTime:     today.Add(-1 * time.Hour),
			ServiceNote: "Reports on and off mild anxiety prior to that. Reports intrusive thoughts of failing classes.",
			Location: Location{
				Address:   "789 Pine Rd, Nowhere, USA",
				Latitude:  34.062235,
				Longitude: -118.233683,
			},
			Status: "completed",
			StartVisit: &VisitTime{
				Timestamp: today.Add(-2 * time.Hour),
				Location: Location{
					Latitude:  34.062235,
					Longitude: -118.233683,
				},
			},
			EndVisit: &VisitTime{
				Timestamp: today.Add(-1 * time.Hour),
				Location: Location{
					Latitude:  34.062235,
					Longitude: -118.233683,
				},
			},
			Tasks: []Task{
				{ID: "3-1", ScheduleID: "3", TaskName: "Check Up", Description: "Check vitals of client and create the report", Completed: true},
				{ID: "3-2", ScheduleID: "3", TaskName: "Bandages", Description: "Change bandages of client wound and clean the wound", Completed: true},
			},
		},
	}
}
