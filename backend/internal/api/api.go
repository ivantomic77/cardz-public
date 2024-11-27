package api

import (
	"log"
	"net/http"
	"time"

	"github.com/FlamingoTP/cardz/internal/handlers"
)

type Application struct {
	Config Config
}

type Config struct {
	Addr string
}

func (app *Application) Mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", handlers.GetHealth)

	return mux
}

func (app *Application) Run(mux *http.ServeMux) error {
	srv := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started at %s", app.Config.Addr)

	return srv.ListenAndServe()
}
