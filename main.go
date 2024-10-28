package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ameliaikeda/tabeo/application"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/danielgtaylor/huma/v2/humacli"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Create a new CLI with Huma. See https://huma.rocks for the docs.
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		router := http.NewServeMux()
		api := humago.New(router, huma.DefaultConfig("SpaceTrouble", "1.0.0"))

		// set up our application.
		app := application.New(database(options))

		// register our application's routes.
		app.Register(api)

		server := http.Server{
			Addr:    fmt.Sprintf(":%d", options.Port),
			Handler: router,
		}

		hooks.OnStart(func() {
			if err := server.ListenAndServe(); err != nil {
				log.Fatalf("http error: %v", err)
			}
		})

		hooks.OnStop(func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				log.Fatalf("http shutdown error: %v", err)
			}

			if err := app.Shutdown(ctx); err != nil {
				log.Fatalf("app shutdown error: %v", err)
			}
		})
	})

	cli.Run()
}

// Options for the CLI. Pass `--port` or set the `SERVICE_PORT` env var.
type Options struct {
	// HTTP Options
	Port int `help:"HTTP Port to listen on" short:"p" default:"8888"`

	DBHost    string `name:"db-host" help:"Database host" default:"localhost"`
	DBPort    int    `name:"db-port" help:"Database port" default:"5432"`
	DBUser    string `name:"db-user" help:"Database username"`
	DBPass    string `name:"db-pass" help:"Database password"`
	DBSSLMode string `name:"db-ssl-mode" help:"Database SSLMode"`
	DBName    string `name:"db-name" help:"Database name"`
	DBDriver  string `name:"db-driver" help:"Database driver name (if different)"`
}

func database(opts *Options) *gorm.DB {
	cfg := postgres.Config{}

	parts := []string{
		fmt.Sprintf("host=%s", opts.DBHost),
		fmt.Sprintf("port=%d", opts.DBPort),
	}

	if opts.DBUser != "" {
		parts = append(parts, fmt.Sprintf("user=%s", opts.DBUser))
	}

	if opts.DBName != "" {
		parts = append(parts, fmt.Sprintf("dbname=%s", opts.DBName))
	}

	if opts.DBPass != "" {
		parts = append(parts, fmt.Sprintf("password=%s", opts.DBPass))
	}

	if opts.DBSSLMode != "" {
		parts = append(parts, fmt.Sprintf("sslmode=%s", opts.DBSSLMode))
	}

	if opts.DBDriver != "" {
		cfg.DriverName = opts.DBDriver
	}

	cfg.DSN = strings.Join(parts, " ")

	db, err := gorm.Open(postgres.New(cfg))
	if err != nil {
		log.Fatalf("database error: %v", err)
	}

	return db
}
