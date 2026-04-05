package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/fulltimegodev/hotel-reservation/api"
	"github.com/fulltimegodev/hotel-reservation/db"
	"github.com/fulltimegodev/hotel-reservation/db/fixtures"
	"github.com/fulltimegodev/hotel-reservation/ent"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	var (
		ctx   = context.Background()
		dbURL = os.Getenv("DATABASE_URL")
	)

	client, err := ent.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	hotelStore := db.NewEntHotelStore(client)
	store := &db.Store{
		User:    db.NewEntUserStore(client),
		Booking: db.NewEntBookingStore(client),
		Room:    db.NewEntRoomStore(client, hotelStore),
		Hotel:   hotelStore,
	}

	user := fixtures.AddUser(store, "james", "foo", false)
	fmt.Println("james ->", api.CreateTokenFromUser(user))
	admin := fixtures.AddUser(store, "admin", "admin", true)
	fmt.Println("admin ->", api.CreateTokenFromUser(admin))
	hotel := fixtures.AddHotel(store, "some hotel", "bermuda", 5, nil)
	room := fixtures.AddRoom(store, "large", true, 88.44, hotel.ID)
	booking := fixtures.AddBooking(store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 5))
	fmt.Println("booking ->", booking.ID)

	for i := 0; i < 100; i++ {
		name := fmt.Sprintf("random hotel name %d", i)
		location := fmt.Sprintf("location %d", i)
		fixtures.AddHotel(store, name, location, rand.Intn(5)+1, nil)
	}
}
