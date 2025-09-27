package api

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct{}

// func (wh WorkoutHandler) NewWorkoutHandler() *WorkoutHandler {
func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh WorkoutHandler) HandleGetByWorkoutID(w http.ResponseWriter, r *http.Request) {
	paramsWorkoutID := chi.URLParam(r, "id")
	if paramsWorkoutID == "" {
		http.NotFound(w, r)
		return
	}

	// TODO replace with a proper uuid
	workoutID, err := strconv.ParseInt(paramsWorkoutID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "oh hey it worked. this is the workout ID: %d\n", workoutID)
}

func (wh WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-- pretend this actually created a workout -- ")
}
