package launchpad

import (
	"context"
	"net/http"
)

type Launch struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DateUTC     string `json:"date_utc"`
	DateUnix    int64  `json:"date_unix"`
	DateLocal   string `json:"date_local"`
	LaunchpadID string `json:"launchpad"`
}

func (a *api) UpcomingLaunches(ctx context.Context) ([]Launch, error) {
	var launches []Launch

	if err := a.Call(ctx, http.MethodGet, "v5/launches/upcoming", nil, &launches); err != nil {
		return nil, err
	}

	return launches, nil
}
