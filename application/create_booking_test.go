package application

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func exampleRequest() *CreateBookingRequest {
	dob, _ := time.Parse("2006-01-02", "1990-06-01")
	tol, _ := time.Parse("2006-01-02", "2025-06-01")

	return &CreateBookingRequest{
		LaunchpadID: "5e9e4502f509094188566f88",
		FirstName:   "Jane",
		LastName:    "Doe",
		Gender:      "female",
		DateOfBirth: dob,
		LaunchDate:  tol,
		Destination: "mars",
	}
}

func TestApplication_CreateBooking(t *testing.T) {
	app := &Application{Repo: Repo{}}

	request := exampleRequest()

	rsp, err := app.CreateBooking(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, rsp)
	require.NotNil(t, rsp.Booking)

	assert.Equal(t, "test uuid", rsp.Booking.ID)
	assert.Equal(t, request.LaunchpadID, rsp.Booking.LaunchpadID)
	assert.Equal(t, request.FirstName, rsp.Booking.FirstName)
	assert.Equal(t, request.LastName, rsp.Booking.LastName)
	assert.Equal(t, request.Gender, rsp.Booking.Gender)
	assert.Equal(t, request.DateOfBirth, rsp.Booking.DateOfBirth)
	assert.Equal(t, request.LaunchDate, rsp.Booking.LaunchDate)
	assert.Equal(t, request.Destination, rsp.Booking.Destination)
}
