package todo

import (
	"app/src/internal/config"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Todo struct {
	Slug  string `json:"slug"`
	Title string `json: "title"`
	Body  string `json: "body"`
}

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{todoID}", GetATodo(configuration))

	return router
}

func GetATodo(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todoID := chi.URLParam(r, "todoID")
		todos := Todo{
			Slug:  todoID,
			Title: "Hello world" + configuration.Constants.PORT,
			Body:  "Helloooooooo",
		}
		render.JSON(w, r, todos)
	}
}
