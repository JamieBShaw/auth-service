package grpc

import (
	"context"
	"github.com/JamieBShaw/auth-service/protob"
	"github.com/JamieBShaw/auth-service/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcServer struct {
	protob.UnimplementedAuthServiceServer
	service service.AuthService
}

func NewGrpcServer(authService service.AuthService) protob.AuthServiceServer {
	return &grpcServer{
		service: authService,
	}
}

func (gs *grpcServer) CreateAccessToken(ctx context.Context, req *protob.CreateAccessTokenRequest) (*protob.CreateAccessTokenResponse, error) {
	if req == nil || req.GetID() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id in request")
	}

	authTokens, err := gs.service.Create(req.GetID())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to create token, error: %v", err.Error())
	}

	return &protob.CreateAccessTokenResponse{
		AuthToken:    authTokens.AccessToken,
		RefreshToken: authTokens.RefreshToken,
	}, nil
}

func (gs *grpcServer) DeleteAccessToken(ctx context.Context, req *protob.DeleteAccessTokenRequest) (*protob.DeleteAccessTokenResponse, error) {
	if req == nil || req.GetAccessUuid() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request")
	}

	err := gs.service.Delete(req.GetAccessUuid())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to delete token, error: %v", err.Error())
	}


	return &protob.DeleteAccessTokenResponse{
		Confirmation: "User access token deleted",
	}, nil
}
