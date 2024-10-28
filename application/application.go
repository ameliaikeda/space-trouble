package application

import (
	"context"
	"errors"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

type Application struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Application {
	return &Application{
		DB: db,
	}
}

func (a *Application) Shutdown(ctx context.Context) error {
	db, err := a.DB.DB()
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
	//
}
