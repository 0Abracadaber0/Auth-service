package validation

import (
	"regexp"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	emptyValue = 0
)

func validateEmail(email string) error {
	if email == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		return status.Error(codes.InvalidArgument, "wrong email format")
	}

	return nil
}

func validatePassword(pass string) error {
	if pass == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}

	if len(pass) < 8 {
		return status.Error(codes.InvalidArgument, "password must be at least 8 characters long")
	}
	return nil
}

func validateAppId(appId int32) error {
	if appId == emptyValue {
		return status.Error(codes.InvalidArgument, "appId is required")
	}
	return nil
}
