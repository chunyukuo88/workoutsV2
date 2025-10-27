package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"

	"github.com/chunyukuo88/workoutsV2/internal/store"
	"github.com/chunyukuo88/workoutsV2/internal/utils"
)

type UserHandler struct {
	userStore store.UserStore
	logger    *log.Logger
}

type registerUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

func (u *UserHandler) validateRegisterRequest(req *registerUserRequest) error {
	if req.Username == "" {
		return errors.New("username is required")
	}
	if len(req.Username) > 50 {
		return errors.New("username is longer than 50 characters")
	}
	if req.Email == "" {
		return errors.New("email is required")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(req.Email) {
		return errors.New("invalid email format")
	}

	if req.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func NewUserHandler(userStore store.UserStore, logger *log.Logger) *UserHandler {
	return &UserHandler{
		userStore,
		logger,
	}
}

func (uh UserHandler) HandleGetUserByID(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.ReadIDParam(r)
	if err != nil {
		uh.logger.Printf("ERROR: failed to read ID param: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid workout ID"})
		return
	}

	user, err := uh.userStore.GetUserByID(int(userID))
	if err != nil {
		uh.logger.Printf("ERROR: failed to get user: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "internal service error"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"user": user})
}

func (uh UserHandler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	var req registerUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		uh.logger.Printf("ERROR: decoding create user: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request"})
		return
	}

	user := &store.User{
		Username: req.Username,
		Email:    req.Email,
	}
	if req.Bio != "" {
		user.Bio = req.Bio
	}

	err = user.PasswordHash.Set(req.Password)
	if err != nil {
		uh.logger.Printf("ERROR: hashing passwords %v:", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	createdUser, err := uh.userStore.CreateUser(user)
	if err != nil {
		uh.logger.Printf("ERROR: createUser: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "failed to create user"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"user": createdUser})
}
