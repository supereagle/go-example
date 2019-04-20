package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var srv http.Server
	cancel := func() {
		// http.Server.Shutdown() is supported in Go v1.8+.
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Fatalf("HTTP server shutdowns with err: %v", err)
		}

		log.Println("HTTP server graceful shutdowns!")
	}
	go GracefulShutdown(cancel)

	log.Println("Starting HTTP server!")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server failed to start: %v", err)
	}
}

// once ensures GracefulShutdown is called at most once. When GracefulShutdown is called,
// this channel would be closed, if another call to close it again, panic would occur.
var once = make(chan struct{})

// GracefulShutdown catches signals of Interrupt, SIGINT, SIGTERM, SIGQUIT and cancel a context.
// If any signals caught, it will call the CancelFunc to cancel a context.
func GracefulShutdown(cancel context.CancelFunc) {
	close(once)

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-c
	log.Printf("Catch signal: %s", s)
	if cancel != nil {
		cancel()
	}

	os.Exit(0)
}
