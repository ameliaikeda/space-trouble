package application

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestApplication_DeleteBooking(t *testing.T) {
	app := &Application{Repo: Repo{}}

	request := &DeleteBookingRequest{ID: "test_id"}

	rsp, err := app.DeleteBooking(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, rsp)

	assert.True(t, rsp.Deleted)
}
