package models

import (
	"time"
)

type EventInformation struct {
	ID          string    `json:"event_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Location    string    `json:"location"`
	TimeZone    string    `json:"time_zone"`
	URL         string    `json:"url"`
	OrganizerID int       `json:"-"`
}

type DisplayEventInfo struct {
	ID              string `json:"event_id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	StartTimeString string `json:"start_time_string"`
	EndTimeString   string `json:"end_time_string"`
	Location        string `json:"location"`
	TimeZone        string `json:"time_zone"`
	Status          string `json:"status"` // UPCOMING, OPEN, ENDED
	URL             string `json:"url"`
	OrganizerID     int    `json:"-"`
}
