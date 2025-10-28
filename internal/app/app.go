package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chunyukuo88/workoutsV2/internal/api"
	"github.com/chunyukuo88/workoutsV2/internal/store"
	"github.com/chunyukuo88/workoutsV2/migrations"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	UserHandler    *api.UserHandler
	TokenHandler   *api.TokenHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}
	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	userStore := store.NewPostgresUserStore(pgDB)
	workoutStore := store.NewPostgresWorkoutStore(pgDB)
	tokenStore := store.NewPostgresTokenStore(pgDB)

	userHandler := api.NewUserHandler(userStore, logger)
	workoutHandler := api.NewWorkoutHandler(workoutStore, logger)
	tokenHandler := api.NewTokenHandler(tokenStore, userStore, logger)

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		UserHandler:    userHandler,
		TokenHandler:   tokenHandler,
		DB:             pgDB,
	}

	return app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yeah it's available, whaddaya want\n")
}
