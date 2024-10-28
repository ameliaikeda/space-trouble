package application

import (
	"context"
	"fmt"
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

func (a *Application) ValidateLaunchpad(ctx context.Context, id string) error {
	// validate the launchpad actually exists
	pads, err := a.API.Launchpads(ctx)
	if err != nil {
		return huma.Error500InternalServerError("failed to check launchpads", err)
	}

	found := false
	for _, pad := range pads {
		if pad.ID == id {
			found = true
			break
		}
	}

	if !found {
		return huma.Error400BadRequest(fmt.Sprintf("launchpad '%s' does not exist", id))
	}

	return nil
}

func (a *Application) ValidateLaunchDate(ctx context.Context, date time.Time, launchpadID string) error {
	year, month, day := time.Now().Date()
	now := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	if date.Before(now) {
		return huma.Error400BadRequest("launch date cannot be in the past")
	}

	launches, err := a.API.UpcomingLaunches(ctx)
	if err != nil {
		return huma.Error500InternalServerError("failed to check launches", err)
	}

	formatted := date.Format(time.DateOnly)

	for _, launch := range launches {
		// we don't care if the launch is TBD or not - if it's in an api slot, we bail out.
		t := time.Unix(launch.DateUnix, 0).Format(time.DateOnly)

		if formatted == t && launch.LaunchpadID == launchpadID {
			return huma.Error400BadRequest("another launch clashes with that launch date")
		}
	}

	return nil
}

// CreateBooking will only allow a booking if the following are true:
// - The launch date is in the future
// - The launchpad given actually exists
// - There are no launches scheduled for the given launchpad on the launch date.
//
// It does not actually check if there is a booking on a given day here - it could easily be changed to do so.
// Whether we allow more than one person to book onto a flight is a question to ask around requirements gathering.
func (a *Application) CreateBooking(ctx context.Context, req *CreateBookingRequest) (*CreateBookingResponse, error) {
	if err := a.ValidateLaunchpad(ctx, req.LaunchpadID); err != nil {
		return nil, err
	}

	if err := a.ValidateLaunchDate(ctx, req.LaunchDate, req.LaunchpadID); err != nil {
		return nil, err
	}

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
