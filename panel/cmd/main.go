package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"sugoi-api/panel/internal/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cfg, err := config.New(filepath.Join(pwd, "config"))
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/create.html")
	})

	r.Get("/images", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/images.html")
	})

	r.Get("/tags", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/tags.html")
	})

	slog.Info(fmt.Sprintf("Listening on %v port", cfg.App.Port))

	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}
