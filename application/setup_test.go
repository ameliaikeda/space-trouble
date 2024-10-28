package application

import (
	"context"
	"github.com/danielgtaylor/huma/v2"
	"time"

	"github.com/ameliaikeda/tabeo/models"
	"github.com/ameliaikeda/tabeo/repository"
)

type Repo struct {
	//
}

func example(id, first, last string) *models.Booking {
	return &models.Booking{
		ID:          id,
		FirstName:   first,
		LastName:    last,
		Gender:      models.GenderUnspecified,
		DateOfBirth: time.Date(2000, time.July, 1, 0, 0, 0, 0, time.UTC),
		LaunchpadID: "test_id",
		LaunchDate:  time.Date(2024, time.July, 1, 0, 0, 0, 0, time.UTC),
		Destination: models.DestinationGanymede,
	}
}

func exampleWithLaunchpad(id, first, last, launchpad string) *models.Booking {
	b := example(id, first, last)
	b.LaunchpadID = launchpad

	return b
}

func (r Repo) CreateBooking(_ context.Context, booking *models.Booking) (*models.Booking, error) {
	booking.ID = "test uuid"
	return booking, nil
}

func (r Repo) DeleteBooking(_ context.Context, id string) (bool, error) {
	switch id {
	case "invalid":
		return false, nil
	case "error":
		return false, huma.Error500InternalServerError("db error")
	}

	return true, nil
}

func (r Repo) ListBookings(_ context.Context) ([]*models.Booking, error) {
	return []*models.Booking{
		example("0", "Jane", "Doe"),
		example("1", "John", "Doe"),
		example("2", "John", "Deer"),
		example("3", "Deary", "Me"),
		example("4", "Fake", "Name"),
		example("5", "Joe", "Bloggs"),
	}, nil
}

func (r Repo) ListBookingsForLaunchpad(_ context.Context, s string) ([]*models.Booking, error) {
	return []*models.Booking{
		exampleWithLaunchpad("0", "Jane", "Doe", s),
		exampleWithLaunchpad("1", "John", "Doe", s),
		exampleWithLaunchpad("2", "John", "Deer", s),
		exampleWithLaunchpad("3", "Deary", "Me", s),
	}, nil
}

var _ repository.Bookings = (*Repo)(nil)
