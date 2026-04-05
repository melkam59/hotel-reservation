package db

import (
	"context"

	"github.com/fulltimegodev/hotel-reservation/ent"
	"github.com/fulltimegodev/hotel-reservation/ent/room"
	"github.com/fulltimegodev/hotel-reservation/types"
)

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
	GetRooms(context.Context, Map) ([]*types.Room, error)
}

type EntRoomStore struct {
	client *ent.Client

	HotelStore
}

func NewEntRoomStore(client *ent.Client, hotelStore HotelStore) *EntRoomStore {
	return &EntRoomStore{
		client:     client,
		HotelStore: hotelStore,
	}
}

func (s *EntRoomStore) GetRooms(ctx context.Context, filter Map) ([]*types.Room, error) {
	query := s.client.Room.Query()

	if hotelID, ok := filter["hotelID"].(int); ok {
		query = query.Where(room.HotelIDEQ(hotelID))
	}

	rooms, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	var res []*types.Room
	for _, r := range rooms {
		res = append(res, toTypeRoom(r))
	}
	return res, nil
}

func (s *EntRoomStore) InsertRoom(ctx context.Context, roomType *types.Room) (*types.Room, error) {
	r, err := s.client.Room.Create().
		SetSize(roomType.Size).
		SetSeaside(roomType.Seaside).
		SetPrice(roomType.Price).
		SetHotelID(roomType.HotelID).
		Save(ctx)

	if err != nil {
		return nil, err
	}
	roomType.ID = r.ID

	// update the hotel with this room id (though ent schema relationship does this via room insertion already!)

	return roomType, nil
}

func toTypeRoom(r *ent.Room) *types.Room {
	if r == nil {
		return nil
	}
	return &types.Room{
		ID:      r.ID,
		Size:    r.Size,
		Seaside: r.Seaside,
		Price:   r.Price,
		HotelID: r.HotelID,
	}
}
