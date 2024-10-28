package application

import (
	"context"
	"github.com/danielgtaylor/huma/v2"

	"github.com/ameliaikeda/tabeo/models"
)

type ListBookingsRequest struct {
	LaunchpadID string `path:"launchpad_id" doc:"Launchpad ID to filter by" example:"5e9e4502f509094188566f88" required:"false"`
}

type ListBookingsResponse struct {
	Bookings []*models.Booking `json:"bookings" doc:"An array of Bookings"`
}

func (a *Application) ListBookings(ctx context.Context, req *ListBookingsRequest) (*ListBookingsResponse, error) {
	var bookings []*models.Booking
	var err error

	switch {
	case req.LaunchpadID != "":
		bookings, err = a.Repo.ListBookingsForLaunchpad(ctx, req.LaunchpadID)
	default:
		bookings, err = a.Repo.ListBookings(ctx)
	}

	if err != nil {
		return nil, huma.Error500InternalServerError("failed to fetch bookings", err)
	}

	return &ListBookingsResponse{Bookings: bookings}, nil
}
