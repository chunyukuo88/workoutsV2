package api

import (
	"errors"
	"testing"
)

func TestValidateRegisterRequest(t *testing.T) {
	h := &UserHandler{}

	tests := []struct {
		name    string
		req     *registerUserRequest
		wantErr error
	}{
		{
			name:    "missing username",
			req:     &registerUserRequest{Username: ""},
			wantErr: errors.New("username is required"),
		},
		{
			name:    "valid username",
			req:     &registerUserRequest{Username: "alex"},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
