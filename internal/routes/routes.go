package routes

import (
	"github.com/chunyukuo88/workoutsV2/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	// r.Group(func(r chi.Router) {
	r.Route("/workouts", func(r chi.Router) {
		r.Use(app.Middleware.Authenticate)

		r.Post("/", app.Middleware.RequireLoggedIn(app.WorkoutHandler.HandleCreateWorkout))
		r.Get("/{id}", app.Middleware.RequireLoggedIn(app.WorkoutHandler.HandleGetByWorkoutID))
		r.Put("/{id}", app.Middleware.RequireLoggedIn(app.WorkoutHandler.HandleUpdateWorkoutByID))
		r.Delete("/{id}", app.Middleware.RequireLoggedIn(app.WorkoutHandler.HandleDeleteWorkout))
	})

	r.Post("/users", app.UserHandler.HandleRegisterUser)
	r.Get("/health", app.HealthCheck)
	r.Post("/tokens/authentication", app.TokenHandler.HandleCreateToken)
	return r
}
