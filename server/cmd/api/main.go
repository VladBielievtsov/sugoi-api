package main

import (
	"log"
	"net/http"
	"sugoi-api/db"
	"sugoi-api/handlers"
	"sugoi-api/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := db.CreateDatabase(); err != nil {
		log.Fatal(err)
	}

	db.Migrate()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		utils.JSONResponse(w, http.StatusOK, "Sugoi-API")
	})

	r.Group(func(r chi.Router) {
		r.Post("/images", handlers.StoreImage)
		r.Get("/images", handlers.GetImages)
	})

	err := http.ListenAndServe(":4000", r)
	if err != nil {
		return
	}
}
