package types

import (
	"time"
)

type Booking struct {
	ID         int       `json:"id,omitempty"`
	UserID     int       `json:"userID,omitempty"`
	RoomID     int       `json:"roomID,omitempty"`
	NumPersons int       `json:"numPersons,omitempty"`
	FromDate   time.Time `json:"fromDate,omitempty"`
	TillDate   time.Time `json:"tillDate,omitempty"`
	Canceled   bool      `json:"canceled"`
}
