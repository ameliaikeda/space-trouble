package repository

import (
	"context"
	"errors"
	"github.com/ameliaikeda/tabeo/models"
	"gorm.io/gorm"
)

type Bookings interface {
	CreateBooking(context.Context, *models.Booking) (*models.Booking, error)
	DeleteBooking(context.Context, *models.Booking) (bool, error)
	ListBookings(context.Context) ([]*models.Booking, error)
	ListBookingsForLaunchpad(context.Context, string) ([]*models.Booking, error)
}

var (
	ErrCreatingBooking  = errors.New("repository: failed to create booking")
	ErrDeletingBooking  = errors.New("repository: failed to delete booking")
	ErrMissingBookingID = errors.New("repository: missing booking ID")
	ErrListingBookings  = errors.New("repository: failed to list bookings")
)

func New(db *gorm.DB) Bookings {
	return &booking{db: db}
}
