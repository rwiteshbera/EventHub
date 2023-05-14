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
	IsOnline    bool      `json:"is_online"`
	URL         string    `json:"url"`
	OrganizerID int       `json:"-"`
}
