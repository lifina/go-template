package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type Server struct {
	*http.Server
}

func NewServer(addr string, port int64, dbAddr string) (*Server, error) {
	api, err := New(dbAddr)
	if err != nil {
		return nil, err
	}

	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%d", addr, port),
		Handler: api,
	}

	return &Server{&srv}, nil
}

func (srv *Server) Start() {
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("Listening on %s\n", srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Println("Shutting down server... Reason:", sig)

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Println("Server gracefully stopped")
}
