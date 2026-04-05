package types

type Hotel struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Rooms    []int  `json:"rooms"`
	Rating   int    `json:"rating"`
}

type Room struct {
	ID int `json:"id,omitempty"`
	// small, normal, kingsize
	Size    string  `json:"size"`
	Seaside bool    `json:"seaside"`
	Price   float64 `json:"price"`
	HotelID int     `json:"hotelID"`
}
