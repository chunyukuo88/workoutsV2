package routes

import (
	"github.com/chunyukuo88/workoutsV2/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(app.Middleware.Authenticate)

		r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetByWorkoutID)
		r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)
		r.Put("/workouts/{id}", app.WorkoutHandler.HandleUpdateWorkoutByID)
		r.Delete("/workouts/{id}", app.WorkoutHandler.HandleDeleteWorkout)
	})

	r.Post("/users", app.UserHandler.HandleRegisterUser)
	r.Get("/health", app.HealthCheck)
	r.Post("/tokens/authentication", app.TokenHandler.HandleCreateToken)
	return r
}
