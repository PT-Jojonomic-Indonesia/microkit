package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Serve(port string, r http.Handler) {

	s := &http.Server{
		ConnState:      CW.OnStateChange,
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    90 * time.Second,
		WriteTimeout:   90 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	time.Sleep(1 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 175*time.Second)
	defer cancel()

	s.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
