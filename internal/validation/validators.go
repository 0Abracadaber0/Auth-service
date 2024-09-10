package validation

import (
	"regexp"
	"unicode"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	emptyValue int32 = 0
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

	if len(pass) > 128 {
		return status.Error(codes.InvalidArgument, "password must be no more than 128 characters long")
	}

	var hasUpper, hasLower, hasDigit, hasSpace, hasInvalidSymb bool

	for _, symb := range pass {
		switch {
		case unicode.IsUpper(symb):
			hasUpper = true
		case unicode.IsLower(symb):
			hasLower = true
		case unicode.IsDigit(symb):
			hasDigit = true
		case unicode.IsSpace(symb):
			hasSpace = true
		case !unicode.IsLetter(symb) && !unicode.IsDigit(symb):
			if !isValidSpecSymb(symb) {
				hasInvalidSymb = true
			}
		}
	}

	if !hasUpper {
		return status.Error(codes.InvalidArgument, "password must contain at least one uppercase letter")
	}
	if !hasLower {
		return status.Error(codes.InvalidArgument, "password must contain at least one lowercase letter")
	}
	if !hasDigit {
		return status.Error(codes.InvalidArgument, "password must contain at least one digit")
	}
	if hasSpace {
		return status.Error(codes.InvalidArgument, "password cannot contain spaces")
	}
	if hasInvalidSymb {
		return status.Error(codes.InvalidArgument, "password can contain only the listed special characters: ~!?@#$%^&*_-+()[]{}><\\/|\"'.,:;")
	}

	return nil
}

func validateAppId(appId int32) error {
	if appId == emptyValue {
		return status.Error(codes.InvalidArgument, "appId is required")
	}
	return nil
}

func validateUserId(userId int32) error {
	if userId == emptyValue {
		return status.Error(codes.InvalidArgument, "userId is required")
	}
	return nil
}

func isValidSpecSymb(symb rune) bool {
	validSpecSymbs := "~!?@#$%^&*_-+()[]{}><\\/|\"'.,:;"
	for _, validSymb := range validSpecSymbs {
		if symb == validSymb {
			return true
		}
	}
	return false
}
