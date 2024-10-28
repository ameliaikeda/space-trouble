package launchpad

import (
	"context"
	"net/http"
)

// Launchpad holds all data about a physical launchpad.
type Launchpad struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	FullName        string   `json:"full_name"`
	Status          string   `json:"status"`
	Launches        []string `json:"launches"`
	Rockets         []string `json:"rockets"`
	LaunchSuccesses int      `json:"launch_successes"`
	LaunchAttempts  int      `json:"launch_attempts"`
	Latitude        float64  `json:"latitude"`
	Longitude       float64  `json:"longitude"`
	Timezone        string   `json:"timezone"`
	Region          string   `json:"region"`
	Locality        string   `json:"locality"`
}

func (a *api) Launchpads(ctx context.Context) ([]Launchpad, error) {
	var rsp []Launchpad

	if err := a.Call(ctx, http.MethodGet, "v4/launchpads", nil, &rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}
