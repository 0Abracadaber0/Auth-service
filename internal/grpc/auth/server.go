package auth

import (
	"context"
	"sso/internal/validation"

	ssov1 "github.com/0Abracadaber0/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {

	validator := &validation.LoginRequestValidator{Request: req}
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	// TODO: implement login via auth service

	return &ssov1.LoginResponse{Token: "1234"}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	panic("implement me")
}
