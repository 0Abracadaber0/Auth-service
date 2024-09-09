package validation

import ssov1 "github.com/0Abracadaber0/protos/gen/go/sso"

type Validator interface {
	Validate() error
}

type LoginRequestValidator struct {
	Request *ssov1.LoginRequest
}

type RegisterRequestValidator struct {
	Request *ssov1.RegisterRequest
}

type IsAdminRequestValidator struct {
	Request *ssov1.IsAdminRequest
}

func (v *LoginRequestValidator) Validate() error {
	if err := validateEmail(v.Request.GetEmail()); err != nil {
		return err
	}

	if err := validatePassword(v.Request.GetPassword()); err != nil {
		return err
	}

	if err := validateAppId(v.Request.GetAppId()); err != nil {
		return err
	}

	return nil
}

func (v *RegisterRequestValidator) Validate() error {
	if err := validateEmail(v.Request.GetEmail()); err != nil {
		return err
	}

	if err := validatePassword(v.Request.GetPassword()); err != nil {
		return err
	}

	return nil
}

func (v *IsAdminRequestValidator) Validate() error {
	if err := validateUserId(v.Request.GetUserId()); err != nil {
		return err
	}

	return nil
}
