package validation_test

import (
	"testing"

	ssov1 "github.com/0Abracadaber0/protos/gen/go/sso"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"sso/internal/validation"
)

type test struct {
	name        string
	request     *ssov1.LoginRequest
	wantErr     bool
	expectedErr error
}

func TestLoginRequestValidator_Validate(t *testing.T) {
	var tests []test

	tests = append(tests, addEmailTests()...)
	tests = append(tests, addPassTests()...)
	tests = append(tests, addAppIdTests()...)

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

func addEmailTests() []test {
	tests := []test{
		{
			"empty email",
			&ssov1.LoginRequest{
				Email:    "",
				Password: "Qwerty123",
				AppId:    1,
			},
			true,
			status.Error(codes.InvalidArgument, "email is required"),
		},
		{
			"good email",
			&ssov1.LoginRequest{
				Email:    "test@example.com",
				Password: "Qwerty123",
				AppId:    1,
			},
			false,
			nil,
		},
		{
			"wrong email",
			&ssov1.LoginRequest{
				Email:    "test-example.com",
				Password: "Qwerty123",
				AppId:    1,
			},
			true,
			status.Error(codes.InvalidArgument, "wrong email format"),
		},
		{
			"wrong email",
			&ssov1.LoginRequest{
				Email:    "test@examplecom",
				Password: "Qwerty123",
				AppId:    1,
			},
			true,
			status.Error(codes.InvalidArgument, "wrong email format"),
		},
		{
			"wrong email",
			&ssov1.LoginRequest{
				Email:    "test@example,com",
				Password: "Qwerty123",
				AppId:    1,
			},
			true,
			status.Error(codes.InvalidArgument, "wrong email format"),
		},
		{
			"wrong email",
			&ssov1.LoginRequest{
				Email:    "@example.com",
				Password: "Qwerty123",
				AppId:    1,
			},
			true,
			status.Error(codes.InvalidArgument, "wrong email format"),
		},
	}

	return tests
}

func addPassTests() []test {
	tests := []test{
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
			"password miss lowercase char",
			&ssov1.LoginRequest{
				Email:    "test@example.com",
				Password: "QWERTY123",
				AppId:    1,
			},
			true,
			status.Error(codes.InvalidArgument, "password must contain at least one lowercase letter"),
		},
		{
			"password miss uppercase char",
			&ssov1.LoginRequest{
				Email:    "test@example.com",
				Password: "qwerty123",
				AppId:    1,
			},
			true,
			status.Error(codes.InvalidArgument, "password must contain at least one uppercase letter"),
		},
		{
			"password miss digit",
			&ssov1.LoginRequest{
				Email:    "test@example.com",
				Password: "QwertyQwerty",
				AppId:    1,
			},
			true,
			status.Error(codes.InvalidArgument, "password must contain at least one digit"),
		},
		{
			"password contain spaces",
			&ssov1.LoginRequest{
				Email:    "test@example.com",
				Password: "Qwerty 123",
				AppId:    1,
			},
			true,
			status.Error(codes.InvalidArgument, "password cannot contain spaces"),
		},
		{
			"password contain invalid char",
			&ssov1.LoginRequest{
				Email:    "test@example.com",
				Password: "Qwerty123`",
				AppId:    1,
			},
			true,
			status.Error(codes.InvalidArgument, "password can contain only the listed special characters: ~!?@#$%^&*_-+()[]{}><\\/|\"'.,:;"),
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
			"good pass",
			&ssov1.LoginRequest{
				Email:    "tes@example.com",
				Password: "Qwerty123!",
				AppId:    1,
			},
			false,
			nil,
		},
	}

	return tests
}

func addAppIdTests() []test {
	tests := []test{
		{
			"empty appId",
			&ssov1.LoginRequest{
				Email:    "test@example.com",
				Password: "Qwerty123",
				AppId:    0,
			},
			true,
			status.Error(codes.InvalidArgument, "appId is required"),
		},
	}

	return tests
}
