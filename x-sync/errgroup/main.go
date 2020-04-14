package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// WithContext returns a new Group and an associated Context derived from ctx.
	//
	// The derived Context is canceled the first time a function passed to Go
	// returns a non-nil error or the first time Wait returns, whichever occurs
	// first.
	egrp, egCtx := errgroup.WithContext(ctx)

	egrp.Go(func() error {
		return signalHandler(egCtx, cancel)
	})

	egrp.Go(func() error {
		return httpSrv(egCtx)
	})

	// Wait blocks until all function calls from the Go method have returned, then
	// returns the first non-nil error (if any) from them.
	err := egrp.Wait()
	if err != nil {
		switch err {
		case http.ErrServerClosed: // ignore
		case context.Canceled: // ignore
		default:
			log.Fatal(err)
		}
	}

	log.Println("done")
}

func signalHandler(ctx context.Context, cancelFn context.CancelFunc) error {
	sigCh := make(chan os.Signal, 2)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		log.Printf("caught signal %v, initiating graceful shutdown", sig)
		cancelFn() // cancel root context
	}()

	<-ctx.Done()
	return ctx.Err()
}

func httpSrv(ctx context.Context) error {
	srv := http.Server{
		Addr: "localhost:8080",
	}

	// handle graceful shutdown
	go func() {
		<-ctx.Done()
		log.Println("initiating shutdown of http server")
		_ = srv.Shutdown(ctx)
	}()

	log.Println("starting http server")
	return srv.ListenAndServe()
}
