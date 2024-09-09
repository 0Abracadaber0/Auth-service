package validation_test

import (
	"testing"

	ssov1 "github.com/0Abracadaber0/protos/gen/go/sso"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"sso/internal/validation"
)

func TestLoginRequestValidator_Validate(t *testing.T) {
	tests := []struct {
		name        string
		request     *ssov1.LoginRequest
		wantErr     bool
		expectedErr error
	}{
		{
			"empty email",
			&ssov1.LoginRequest{
				Email:    "",
				Password: "qwerty123",
				AppId:    1,
			},
			true,
			status.Error(codes.InvalidArgument, "email is required"),
		},
		{
			"empty password",
			&ssov1.LoginRequest{
				Email:    "test@example.com",
				Password: "",
				AppId:    1,
			},
			true,
			status.Error(codes.InvalidArgument, "password is required"),
		},
		{
			"empty appId",
			&ssov1.LoginRequest{
				Email:    "test@example.com",
				Password: "qwerty123",
				AppId:    0,
			},
			true,
			status.Error(codes.InvalidArgument, "appId is required"),
		},
		{
			"short password",
			&ssov1.LoginRequest{
				Email:    "test@example.com",
				Password: "short",
				AppId:    1,
			},
			true,
			status.Error(codes.InvalidArgument, "password must be at least 8 characters long"),
		},
		{
			"good data",
			&ssov1.LoginRequest{
				Email:    "test@example.com",
				Password: "qwerty123",
				AppId:    1,
			},
			false,
			nil,
		},
		{
			"wrong email",
			&ssov1.LoginRequest{
				Email:    "test-example.com",
				Password: "qwerty123",
				AppId:    1,
			},
			true,
			status.Error(codes.InvalidArgument, "wrong email format"),
		},
		{
			"wrong email",
			&ssov1.LoginRequest{
				Email:    "test@examplecom",
				Password: "qwerty123",
				AppId:    1,
			},
			true,
			status.Error(codes.InvalidArgument, "wrong email format"),
		},
		{
			"wrong email",
			&ssov1.LoginRequest{
				Email:    "test@example,com",
				Password: "qwerty123",
				AppId:    1,
			},
			true,
			status.Error(codes.InvalidArgument, "wrong email format"),
		},
		{
			"wrong email",
			&ssov1.LoginRequest{
				Email:    "@example.com",
				Password: "qwerty123",
				AppId:    1,
			},
			true,
			status.Error(codes.InvalidArgument, "wrong email format"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := &validation.LoginRequestValidator{Request: tt.request}
			err := validator.Validate()
			if tt.wantErr {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
