package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/atuy1213/textintoimage/api/internal"
	"github.com/go-chi/chi/v5"
)

func main() {
	os.Exit(run())
}

func run() int {
	const (
		statusOK = 0
		statusNG = 1
	)

	const (
		readHeaderTimeout = 2
		ctxTimeout        = 5
	)

	ctx := context.Background()

	r := chi.NewMux()
	internal.InitRouter(r)
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: time.Second * readHeaderTimeout,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	go func() {
		slog.Info("start server")
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			slog.Error("failed to ListenAndServe", err)
		}
	}()

	<-ctx.Done()

	ctx, timeout := context.WithTimeout(context.Background(), time.Second*ctxTimeout)
	defer timeout()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("failed to Shutdown", err)
		return statusNG
	}

	return statusOK
}
