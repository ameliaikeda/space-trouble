package application

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestApplication_ListBookings(t *testing.T) {
	app := &Application{Repo: Repo{}}

	request := &ListBookingsRequest{}

	rsp, err := app.ListBookings(context.Background(), request)

	require.NoError(t, err)
	require.NotNil(t, rsp)
	require.NotNil(t, rsp.Bookings)

	assert.Len(t, rsp.Bookings, 6)
}

func TestApplication_ListBookingsWithLaunchpad(t *testing.T) {
	app := &Application{Repo: Repo{}}

	request := &ListBookingsRequest{LaunchpadID: "testing"}

	rsp, err := app.ListBookings(context.Background(), request)

	require.NoError(t, err)
	require.NotNil(t, rsp)
	require.NotNil(t, rsp.Bookings)

	assert.Len(t, rsp.Bookings, 4)

	for i, b := range rsp.Bookings {
		assert.Equal(t, request.LaunchpadID, b.LaunchpadID)
		assert.Equal(t, strconv.Itoa(i), b.ID)
	}
}
