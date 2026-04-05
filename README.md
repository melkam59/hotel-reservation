# Hotel Reservation Backend

A simple, robust hotel reservation API built with Go, Fiber, and PostgreSQL (via Ent ORM).

## Features
- **Users**: Can browse hotels and book rooms.
- **Admins**: Can manage properties and supervise reservations/bookings.
- **Authentication**: JWT-based token authentication and authorization.
- **API**: Full CRUD endpoints for Hotels, Rooms, and Bookings dealing with structured JSON.
- **ORM**: Powered by the [ent](https://entgo.io/) framework for type-safe database schema and migrations.

## Environment Variables

Create a `.env` file in the root of the project with the following configuration:

```env
HTTP_LISTEN_ADDRESS=:3000
JWT_SECRET=somethingsupersecretthatNOBODYKNOWS
DATABASE_URL=postgres://your_user:your_password@localhost/hotel_reservation?sslmode=disable
```

## Getting Started

1. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

2. **Initialize Database Schema & Run Server**:
   ```bash
   go run main.go
   ```

3. **(Optional) Seed the Database**:
   Populates the database with initial users, hotels, and testing resources.
   ```bash
   go run scripts/seed.go
   ```
