package api

import (
	"strconv"

	"github.com/fulltimegodev/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
)

type BookingHandler struct {
	store *db.Store
}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{
		store: store,
	}
}

func (h *BookingHandler) HandleCancelBooking(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrInvalidID()
	}
	booking, err := h.store.Booking.GetBookingByID(c.Context(), id)
	if err != nil {
		return ErrNotResourceNotFound("booking")
	}
	user, err := getAuthUser(c)
	if err != nil {
		return ErrUnAuthorized()
	}
	if booking.UserID != user.ID {
		return ErrUnAuthorized()
	}
	if err := h.store.Booking.UpdateBooking(c.Context(), id, db.Map{"canceled": true}); err != nil {
		return err
	}
	return c.JSON(genericResp{Type: "msg", Msg: "updated"})
}

func (h *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	bookings, err := h.store.Booking.GetBookings(c.Context(), db.Map{})
	if err != nil {
		return ErrNotResourceNotFound("bookings")
	}
	return c.JSON(bookings)
}

func (h *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ErrInvalidID()
	}
	booking, err := h.store.Booking.GetBookingByID(c.Context(), id)
	if err != nil {
		return ErrNotResourceNotFound("booking")
	}
	user, err := getAuthUser(c)
	if err != nil {
		return ErrUnAuthorized()
	}
	if booking.UserID != user.ID {
		return ErrUnAuthorized()
	}
	return c.JSON(booking)
}
