package db

import (
	"context"

	"github.com/fulltimegodev/hotel-reservation/ent"
	"github.com/fulltimegodev/hotel-reservation/ent/hotel"
	"github.com/fulltimegodev/hotel-reservation/types"
)

type HotelStore interface {
	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)
	Update(context.Context, Map, Map) error
	GetHotels(context.Context, Map, *Pagination) ([]*types.Hotel, error)
	GetHotelByID(context.Context, int) (*types.Hotel, error)
}

type EntHotelStore struct {
	client *ent.Client
}

func NewEntHotelStore(client *ent.Client) *EntHotelStore {
	return &EntHotelStore{
		client: client,
	}
}

func (s *EntHotelStore) GetHotelByID(ctx context.Context, id int) (*types.Hotel, error) {
	h, err := s.client.Hotel.Query().Where(hotel.IDEQ(id)).WithRooms().Only(ctx)
	if err != nil {
		return nil, err
	}
	return toTypeHotel(h), nil
}

func (s *EntHotelStore) GetHotels(ctx context.Context, filter Map, pag *Pagination) ([]*types.Hotel, error) {
	query := s.client.Hotel.Query().WithRooms()
	if rating, ok := filter["rating"].(int); ok && rating > 0 {
		query = query.Where(hotel.RatingEQ(rating))
	}

	if pag != nil {
		query = query.Offset(int((pag.Page - 1) * pag.Limit)).Limit(int(pag.Limit))
	}

	hotels, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	var res []*types.Hotel
	for _, h := range hotels {
		res = append(res, toTypeHotel(h))
	}
	return res, nil
}

func (s *EntHotelStore) Update(ctx context.Context, filter Map, update Map) error {
	id, ok := filter["_id"].(int)
	if !ok {
		return nil // Simplification for this example
	}
	upd := s.client.Hotel.UpdateOneID(id)

	if rooms, ok := update["$push"].(map[string]any); ok {
		if roomID, ok := rooms["rooms"].(int); ok {
			upd.AddRoomIDs(roomID)
		}
	}

	_, err := upd.Save(ctx)
	return err
}

func (s *EntHotelStore) InsertHotel(ctx context.Context, hotelType *types.Hotel) (*types.Hotel, error) {
	h, err := s.client.Hotel.Create().
		SetName(hotelType.Name).
		SetLocation(hotelType.Location).
		SetRating(hotelType.Rating).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	hotelType.ID = h.ID
	return hotelType, nil
}

func toTypeHotel(h *ent.Hotel) *types.Hotel {
	if h == nil {
		return nil
	}
	res := &types.Hotel{
		ID:       h.ID,
		Name:     h.Name,
		Location: h.Location,
		Rating:   h.Rating,
	}
	if h.Edges.Rooms != nil {
		for _, r := range h.Edges.Rooms {
			res.Rooms = append(res.Rooms, r.ID)
		}
	}
	return res
}
