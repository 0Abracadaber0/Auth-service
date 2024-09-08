package validation

import ssov1 "github.com/0Abracadaber0/protos/gen/go/sso"

type Validator interface {
	Validate() error
}

type LoginRequestValidator struct {
	Request *ssov1.LoginRequest
}

func (v *LoginRequestValidator) Validate() error {
	if err := validateEmail(v.Request.Email); err != nil {
		return err
	}

	if err := validatePassword(v.Request.Password); err != nil {
		return err
	}

	if err := validateAppId(v.Request.AppId); err != nil {
		return err
	}

	return nil
}
