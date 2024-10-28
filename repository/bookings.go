package repository

import (
	"context"
	"errors"
	"github.com/ameliaikeda/tabeo/lib/uuid"
	"github.com/ameliaikeda/tabeo/models"
	"gorm.io/gorm"
)

var _ Bookings = (*booking)(nil)

type booking struct {
	db *gorm.DB
}

func (r *booking) CreateBooking(ctx context.Context, booking *models.Booking) (*models.Booking, error) {
	if booking.ID == "" {
		booking.ID = uuid.NewString()
	}

	if err := r.db.WithContext(ctx).Create(&booking).Error; err != nil {
		return nil, errors.Join(ErrCreatingBooking, err)
	}

	return booking, nil
}

func (r *booking) DeleteBooking(ctx context.Context, booking *models.Booking) (bool, error) {
	b := &models.Booking{ID: booking.ID}

	if b.ID == "" {
		return false, ErrMissingBookingID
	}

	tx := r.db.WithContext(ctx).Delete(b)
	if err := tx.Error; err != nil {
		return false, errors.Join(ErrDeletingBooking, err)
	}

	return tx.RowsAffected == 1, nil
}

func (r *booking) ListBookings(ctx context.Context) ([]*models.Booking, error) {
	bookings := make([]*models.Booking, 0)

	tx := r.db.WithContext(ctx).Order("id asc").Find(&bookings)
	if err := tx.Error; err != nil {
		return nil, errors.Join(ErrListingBookings, err)
	}

	return bookings, nil
}

func (r *booking) ListBookingsForLaunchpad(ctx context.Context, launchpadID string) ([]*models.Booking, error) {
	var bookings []*models.Booking

	tx := r.db.WithContext(ctx).Order("id asc").Where(&models.Booking{LaunchpadID: launchpadID}).Find(&bookings)
	if err := tx.Error; err != nil {
		return nil, errors.Join(ErrListingBookings, err)
	}

	return bookings, nil
}
