package application

import (
	"context"
	"github.com/danielgtaylor/huma/v2"

	"github.com/ameliaikeda/tabeo/models"
)

type ListLaunchpadBookingsRequest struct {
	LaunchpadID string `path:"launchpad_id" doc:"Launchpad ID to filter by" example:"5e9e4502f509094188566f88" required:"false"`
}

type ListBookingsRequest struct {
	//
}

type ListBookingsResponse struct {
	Bookings []*models.Booking `json:"bookings" doc:"An array of Bookings"`
}

func (a *Application) ListBookings(ctx context.Context, _ *ListBookingsRequest) (*ListBookingsResponse, error) {
	bookings, err := a.Repo.ListBookings(ctx)
	if err != nil {
		return nil, huma.Error500InternalServerError("failed to fetch bookings", err)
	}

	return &ListBookingsResponse{Bookings: bookings}, nil
}

func (a *Application) ListBookingsByLaunchpad(ctx context.Context, req *ListLaunchpadBookingsRequest) (*ListBookingsResponse, error) {
	bookings, err := a.Repo.ListBookingsForLaunchpad(ctx, req.LaunchpadID)
	if err != nil {
		return nil, huma.Error500InternalServerError("failed to fetch bookings", err)
	}

	return &ListBookingsResponse{Bookings: bookings}, nil
}
