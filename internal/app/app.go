package app

import (
	"fmt"
	"log"
	"net/http"
	"stroy-svaya/internal/config"
	"stroy-svaya/internal/handler"
	"stroy-svaya/internal/repository"
	"stroy-svaya/internal/service"
	"time"
)

type App struct {
	c   *config.Config
	r   *repository.SQLiteRepository
	s   *service.Service
	h   *handler.Handler
	srv *http.Server
}

func New() (*App, error) {
	a := &App{}
	a.c = config.Load()
	var err error
	a.r, err = repository.NewSQLiteRepository(a.c.DatabasePath)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}
	a.s = service.NewService(a.r)
	a.h = handler.NewHandler(a.s)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /insertpdrline", a.h.InsertPileDrivingRecordLine)
	mux.HandleFunc("GET /getpdrlines", a.h.GetPileDrivingRecord)
	a.srv = &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		IdleTimeout:  time.Second * 120,
		ReadTimeout:  time.Second * 1,
		WriteTimeout: time.Second * 1,
	}
	return a, nil
}

func (a *App) Run() error {
	fmt.Println("Stroy-svaya is running..")
	err := a.srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
