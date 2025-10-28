package api

import (
	"errors"
	"testing"
)

func TestValidateRegisterRequest(t *testing.T) {
	h := &UserHandler{}

	tests := []struct {
		scenario string
		req      *registerUserRequest
		wantErr  error
	}{
		{
			scenario: "missing username",
			req:      &registerUserRequest{Username: ""},
			wantErr:  errors.New("username is required"),
		},
		{
			scenario: "username longer than 50 characters",
			req:      &registerUserRequest{Username: "MorrisBorisAndTheFiveMagicPigsAndTommySalamiAndHamletYogurt"},
			wantErr:  errors.New("username is longer than 50 characters"),
		},
		{
			scenario: "valid username",
			req:      &registerUserRequest{Username: "alex", Email: "woobler@otherhouse.com", Password: "lamepassword"},
			wantErr:  nil,
		},
		{
			scenario: "missing email",
			req:      &registerUserRequest{Username: "woobler", Email: ""},
			wantErr:  errors.New("email is required"),
		},
		{
			scenario: "malformed email",
			req:      &registerUserRequest{Username: "the_fourth", Email: "foo"},
			wantErr:  errors.New("invalid email format"),
		},
		{
			scenario: "missing password",
			req:      &registerUserRequest{Username: "monster", Email: "test@example.com"},
			wantErr:  errors.New("password is required"),
		},
		{
			scenario: "all entries valid",
			req:      &registerUserRequest{Username: "alex", Email: "alex@gochenour.com", Password: "verySecurePassword"},
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			err := h.validateRegisterRequest(tt.req)
			if (err == nil) != (tt.wantErr == nil) {
				t.Fatalf("expected error = %v, got %v", tt.wantErr, err)
			}
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Fatalf("expected error message %q, got %q", tt.wantErr.Error(), err.Error())
			}
		})
	}
}
