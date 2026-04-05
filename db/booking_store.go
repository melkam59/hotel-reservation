package db

import (
	"context"
	"time"

	"github.com/fulltimegodev/hotel-reservation/ent"
	"github.com/fulltimegodev/hotel-reservation/ent/booking"
	"github.com/fulltimegodev/hotel-reservation/types"
)

type BookingStore interface {
	InsertBooking(context.Context, *types.Booking) (*types.Booking, error)
	GetBookings(context.Context, Map) ([]*types.Booking, error)
	GetBookingByID(context.Context, int) (*types.Booking, error)
	UpdateBooking(context.Context, int, Map) error
}

type EntBookingStore struct {
	client *ent.Client
}

func NewEntBookingStore(client *ent.Client) *EntBookingStore {
	return &EntBookingStore{
		client: client,
	}
}

func (s *EntBookingStore) UpdateBooking(ctx context.Context, id int, update Map) error {
	upd := s.client.Booking.UpdateOneID(id)

	if canceled, ok := update["canceled"].(bool); ok {
		upd.SetCanceled(canceled)
	}

	_, err := upd.Save(ctx)
	return err
}

func (s *EntBookingStore) GetBookings(ctx context.Context, filter Map) ([]*types.Booking, error) {
	query := s.client.Booking.Query()

	if roomID, ok := filter["roomID"].(int); ok {
		query = query.Where(booking.RoomIDEQ(roomID))
	}

	if userID, ok := filter["userID"].(int); ok {
		query = query.Where(booking.UserIDEQ(userID))
	}

	if fromDate, ok := filter["fromDate"].(map[string]any); ok {
		if gte, ok := fromDate["$gte"].(time.Time); ok {
			query = query.Where(booking.FromDateGTE(gte))
		}
	}

	if tillDate, ok := filter["tillDate"].(map[string]any); ok {
		if lte, ok := tillDate["$lte"].(time.Time); ok {
			query = query.Where(booking.TillDateLTE(lte))
		}
	}

	bookings, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	var res []*types.Booking
	for _, b := range bookings {
		res = append(res, toTypeBooking(b))
	}
	return res, nil
}

func (s *EntBookingStore) GetBookingByID(ctx context.Context, id int) (*types.Booking, error) {
	b, err := s.client.Booking.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return toTypeBooking(b), nil
}

func (s *EntBookingStore) InsertBooking(ctx context.Context, bookingType *types.Booking) (*types.Booking, error) {
	b, err := s.client.Booking.Create().
		SetUserID(bookingType.UserID).
		SetRoomID(bookingType.RoomID).
		SetNumPersons(bookingType.NumPersons).
		SetFromDate(bookingType.FromDate).
		SetTillDate(bookingType.TillDate).
		Save(ctx)

	if err != nil {
		return nil, err
	}
	bookingType.ID = b.ID
	return bookingType, nil
}

func toTypeBooking(b *ent.Booking) *types.Booking {
	if b == nil {
		return nil
	}
	return &types.Booking{
		ID:         b.ID,
		UserID:     b.UserID,
		RoomID:     b.RoomID,
		NumPersons: b.NumPersons,
		FromDate:   b.FromDate,
		TillDate:   b.TillDate,
		Canceled:   b.Canceled,
	}
}
