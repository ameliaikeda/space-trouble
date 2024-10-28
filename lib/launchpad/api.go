// Package launchpad is a *very* quick and dirty implementation of an API client for the SpaceX API.
// This can and should be done more completely - it has very useful features like searching for all launches by launchpad ID.
// (which is exactly what we'd want), but I'd rather get this finished for now than throw a ton of time into building a proper API client.
package launchpad

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type API interface {
	Launchpads(ctx context.Context) ([]Launchpad, error)
	UpcomingLaunches(ctx context.Context) ([]Launch, error)
}

type api struct {
	BaseURL string // https://api.spacexdata.com
	Client  *http.Client
}

func NewAPI(baseURL string, httpClient *http.Client) API {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &api{
		BaseURL: baseURL,
		Client:  httpClient,
	}
}

func (a *api) Call(ctx context.Context, method, endpoint string, body io.Reader, value any) error {
	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", a.BaseURL, endpoint), body)
	if err != nil {
		return huma.Error500InternalServerError("error calling spacex api", err)
	}

	rsp, err := a.Client.Do(req)
	if err != nil {
		return huma.Error500InternalServerError("error calling spacex api", err)
	}

	defer rsp.Body.Close()

	if err := json.NewDecoder(rsp.Body).Decode(&value); err != nil {
		return huma.Error500InternalServerError("error decoding spacex response", err)
	}

	return nil
}
