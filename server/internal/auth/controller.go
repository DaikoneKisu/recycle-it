package auth

import (
	"context"
	"log/slog"

	pb "github.com/DaikoneKisu/recycle-it/server/internal/protos/auth"
	"gorm.io/gorm"
)

type Controller struct {
	pb.UnimplementedAuthControllerServer
	authenticator Authenticator
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{
		authenticator: NewAuthenticator(db),
	}
}

func (c Controller) SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	jwt, err := c.authenticator.SignUp(ctx, request.Nickname, request.Password)
	if err != nil {
		slog.Debug(err.Error())
		return nil, err
	}
	return &pb.SignUpResponse{Token: jwt}, nil
}

func (c Controller) SignIn(ctx context.Context, request *pb.SignInRequest) (*pb.SignInResponse, error) {
	jwt, err := c.authenticator.SignIn(ctx, request.Nickname, request.Password)
	if err != nil {
		slog.Debug(err.Error())
		return nil, err
	}
	return &pb.SignInResponse{Token: jwt}, nil
}
