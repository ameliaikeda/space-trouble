package repository

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ameliaikeda/tabeo/models"
	_ "github.com/jackc/pgx/v5"
	"github.com/peterldowns/pgtestdb"
	"github.com/peterldowns/pgtestdb/migrators/goosemigrator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB[T testing.TB](t T) *sql.DB {
	t.Helper()

	cfg := pgtestdb.Config{
		DriverName: "pgx",
		Host:       "localhost",
		Port:       "5432",
		User:       os.Getenv("USER"),
		Options:    "sslmode=disable",
	}

	if v := os.Getenv("SERVICE_DB_USER"); v != "" {
		cfg.User = v
	}

	if v := os.Getenv("SERVICE_DB_PASS"); v != "" {
		cfg.Password = v
	}

	if v := os.Getenv("SERVICE_DB_HOST"); v != "" {
		cfg.Host = v
	}

	if v := os.Getenv("SERVICE_DB_PORT"); v != "" {
		cfg.Port = v
	}

	if v := os.Getenv("SERVICE_DB_SSL_MODE"); v != "" {
		cfg.Options = fmt.Sprintf("sslmode=%s", v)
	}

	m := goosemigrator.New("migrations",
		goosemigrator.WithFS(os.DirFS("..")),
		goosemigrator.WithTableName("goose_db_version"))

	return pgtestdb.New(t, cfg, m)
}

func AppContext[T testing.TB](t T) context.Context {
	t.Helper()

	return context.Background()
}

func TestBooking_CreateBooking(t *testing.T) {
	t.Parallel()

	db := NewDB(t)
	gdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	defer func() {
		d, _ := gdb.DB()
		if d != nil {
			_ = d.Close()
		}
	}()

	require.NoError(t, err, "gorm should initialize")

	bk := New(gdb)

	subject := &models.Booking{
		LaunchpadID: "testing ID",
		Destination: models.DestinationMars,
		FirstName:   "Jane",
		LastName:    "Doe",
		Gender:      models.GenderUnspecified,
		DateOfBirth: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		LaunchDate:  time.Date(2024, time.October, 29, 0, 0, 0, 0, time.UTC),
	}

	b, err := bk.CreateBooking(AppContext(t), subject)

	assert.NoError(t, err)
	require.NotEmpty(t, b)

	assert.NotEmpty(t, b.ID)

	assert.Equal(t, subject.LaunchpadID, b.LaunchpadID)
	assert.Equal(t, subject.Destination, b.Destination)
	assert.Equal(t, subject.FirstName, b.FirstName)
	assert.Equal(t, subject.LastName, b.LastName)
	assert.Equal(t, subject.Gender, b.Gender)
	assert.Equal(t, subject.LaunchDate, b.LaunchDate)
	assert.Equal(t, subject.DateOfBirth, b.DateOfBirth)

	bookings, err := bk.ListBookings(AppContext(t))
	require.NoError(t, err)

	assert.Len(t, bookings, 1)
	assert.EqualExportedValues(t, bookings[0], b)
}

func TestBooking_ListBookingsForLaunchpad(t *testing.T) {
	t.Parallel()

	db := NewDB(t)
	gdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	defer func() {
		d, _ := gdb.DB()
		if d != nil {
			_ = d.Close()
		}
	}()

	require.NoError(t, err, "gorm should initialize")

	bk := New(gdb)

	subject := example()

	_, _ = bk.CreateBooking(AppContext(t), example())
	_, _ = bk.CreateBooking(AppContext(t), example())
	_, _ = bk.CreateBooking(AppContext(t), example())

	correct, err := bk.ListBookingsForLaunchpad(AppContext(t), subject.LaunchpadID)
	require.NoError(t, err)

	assert.Len(t, correct, 3)

	incorrect, err := bk.ListBookingsForLaunchpad(AppContext(t), subject.LaunchpadID+"_incorrect")
	require.NoError(t, err)

	assert.Len(t, incorrect, 0)
}

func example() *models.Booking {
	return &models.Booking{
		LaunchpadID: "testing ID",
		Destination: models.DestinationMars,
		FirstName:   "Jane",
		LastName:    "Doe",
		Gender:      models.GenderUnspecified,
		DateOfBirth: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		LaunchDate:  time.Date(2024, time.October, 29, 0, 0, 0, 0, time.UTC),
	}
}
