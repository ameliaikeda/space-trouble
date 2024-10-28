package application

import (
	"context"
	"github.com/danielgtaylor/huma/v2"
)

type DeleteBookingRequest struct {
	ID string `json:"id" doc:"Booking ID to delete" example:"01904d33-d262-7531-b71c-05555c63df91"`
}

type DeleteBookingResponse struct {
	Deleted bool `json:"deleted" doc:"Confirmation of deletion"`
}

func (a *Application) DeleteBooking(ctx context.Context, req *DeleteBookingRequest) (*DeleteBookingResponse, error) {
	ok, err := a.Repo.DeleteBooking(ctx, req.ID)
	if err != nil {
		return nil, huma.Error500InternalServerError("failed to delete booking", err)
	}

	return &DeleteBookingResponse{Deleted: ok}, nil
}
