package auth

import (
	"context"
	"sso/internal/validation"

	ssov1 "github.com/0Abracadaber0/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(ctx context.Context,
		email string,
		password string,
		appId int32,
	) (token string, err error)
	RegisterNewUser(ctx context.Context,
		email string,
		password string,
	) (userId int64, err error)
	IsAdmin(ctx context.Context, userId int64) (bool, error)
}

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {

	validator := &validation.LoginRequestValidator{Request: req}
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetEmail(), req.GetAppId())
	if err != nil {
		// TODO: ...
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssov1.LoginResponse{
		Token: token,
	}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	validator := validation.RegisterRequestValidator(validation.RegisterRequestValidator{Request: req})
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	userId, err := s.auth.RegisterNewUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		// TODO: ...
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssov1.RegisterResponse{
		UserId: userId,
	}, nil
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
	validator := validation.IsAdminRequestValidator{Request: req}
	if err := validator.Validate(); err != nil {
		return nil, err
	}

	isAdmin, err := s.auth.IsAdmin(ctx, int64(req.GetUserId()))
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssov1.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil
}
