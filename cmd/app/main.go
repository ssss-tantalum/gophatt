package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ssss-tantalum/gophatt/pkg/app"
	"github.com/ssss-tantalum/gophatt/pkg/server"
)

func main() {
	cfg, err := app.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	client := app.NewClient(cfg.DbDSN)
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}

	app := app.New(ctx, cfg, client)
	srv := server.New(app)

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	port := fmt.Sprintf(":%d", cfg.HTTPPort)
	go func() {
		if err := srv.Start(port); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
