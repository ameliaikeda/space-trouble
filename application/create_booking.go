package application

import (
	"context"
	"time"

	"github.com/ameliaikeda/tabeo/models"
	"github.com/danielgtaylor/huma/v2"
)

type CreateBookingRequest struct {
	LaunchpadID string    `json:"launchpad_id" doc:"Launchpad ID" example:"5e9e4502f509094188566f88"`
	FirstName   string    `json:"first_name" doc:"First Name" example:"Jane"`
	LastName    string    `json:"last_name" doc:"Last Name" example:"Doe"`
	Gender      string    `json:"gender" doc:"Gender" example:"female" enum:"female,male,unspecified"`
	DateOfBirth time.Time `json:"date_of_birth" doc:"Birthday" example:"1990-01-01" format:"date"`
	LaunchDate  time.Time `json:"launch_date" doc:"Launch Date" example:"2024-01-01" format:"date"`
	Destination string    `json:"destination" doc:"Destination of flight" example:"mars" enum:"mars,moon,pluto,asteroid_belt,europa,titan,ganymede"`
}

type CreateBookingResponse struct {
	Booking *models.Booking `json:"booking" doc:"The created booking"`
}

func (a *Application) CreateBooking(ctx context.Context, req *CreateBookingRequest) (*CreateBookingResponse, error) {
	booking, err := a.Repo.CreateBooking(ctx, &models.Booking{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Gender:      req.Gender,
		DateOfBirth: req.DateOfBirth,
		LaunchpadID: req.LaunchpadID,
		LaunchDate:  req.LaunchDate,
		Destination: req.Destination,
	})

	if err != nil {
		return nil, huma.Error500InternalServerError("failed to create booking", err)
	}

	return &CreateBookingResponse{Booking: booking}, nil
}
