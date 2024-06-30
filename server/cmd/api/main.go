package main

import (
	"log"
	"net/http"
	"sugoi-api/db"
	"sugoi-api/handlers"
	"sugoi-api/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
	}))

	fs := http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads")))
	r.Handle("/uploads/*", fs)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		utils.JSONResponse(w, http.StatusOK, "Sugoi-API")
	})

	r.Group(func(r chi.Router) {
		r.Post("/images", handlers.StoreImage) // Form-Data{ source: string, image: file }
		r.Get("/images", handlers.GetImages)   // Query{ ?page=1, ?page_size=5 }
		r.Get("/images/{id}", handlers.GetImageByID)
		r.Get("/images/random", handlers.GetRandomImages) // Body{ limit: int }
	})

	r.Group(func(r chi.Router) {
		r.Post("/tags", handlers.CreateTag) // Body{ name: string, description: string }
		r.Get("/tags/{name}", handlers.GetTagByName)
		r.Get("/tags", handlers.GetTags)
	})

	r.Group(func(r chi.Router) {
		r.Post("/characters", handlers.CreateCharacter)
		r.Get("/characters", handlers.GetCharacters) // Query { name }
	})

	err := http.ListenAndServe(":4000", r)
	if err != nil {
		return
	}
}
