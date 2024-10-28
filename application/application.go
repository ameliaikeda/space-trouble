package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/ameliaikeda/tabeo/lib/launchpad"
	"github.com/ameliaikeda/tabeo/repository"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

type Application struct {
	Repo repository.Bookings
	API  launchpad.API
	db   *gorm.DB
}

func New(db *gorm.DB) *Application {
	return &Application{
		Repo: repository.New(db),
		db:   db,
	}
}

func (a *Application) Shutdown(ctx context.Context) error {
	db, err := a.db.DB()
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidDB) {
			// we've already been shut down, ignore this
			return nil
		}

		return err
	}

	done := make(chan error)

	go func() {
		done <- db.Close()
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("error shutting down database: %w", ctx.Err())
	case err := <-done:
		return err
	}
}

func (a *Application) Register(api huma.API) {
	huma.Post(api, "/bookings", a.CreateBooking)
	huma.Get(api, "/bookings/{launchpad_id}", a.ListBookings)
	huma.Delete(api, "/bookings", a.DeleteBooking)
}
