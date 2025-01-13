package store

import (
	"context"

	"github.com/dvnthn168/TaskFlow/gen/userpb"
	"github.com/dvnthn168/TaskFlow/pkg/helpers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	repo UserRepository
	userpb.UnimplementedUsersServer
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (h *UserService) CreateUser(ctx context.Context, req *userpb.PayloadWithSignleUser) (*userpb.PayloadWithUserID, error) {

	if req == nil || req.User == nil {
		return nil, status.Error(codes.InvalidArgument, "Request or user information is missing")
	}

	dto := User{
		Username:  req.User.GetUsername(),
		Email:     req.User.GetEmail(),
		Password:  req.User.GetPassword(),
	}

	userID, err := h.repo.CreateUser(ctx, &dto)
	if err != nil {
		return nil, err
	}
	return &userpb.PayloadWithUserID{UserId: userID}, nil
}

// func (h *UserService) VerifyUser(ctx context.Context, req *userpb.VerifyUserRequest) (*emptypb.Empty, error) {
// 	dto := &usermd.VerifyUserDTO{
// 		UserID: req.UserId,
// 		OTP:    req.Otp,
// 	}

// 	err := h.repo.VerifyUser(ctx, dto)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &emptypb.Empty{}, nil
// }

// func (h *UserService) ResendVerifyUser(ctx context.Context, req *userpb.ResendUserRequest) (*emptypb.Empty, error) {

// 	err := h.repo.ResendVerifyUser(ctx, req.UserId, req.GetEmail())
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &emptypb.Empty{}, nil
// }

// func (s *UserService) LoginUser(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {

// 	user, err := s.repo.GetUserByEmail(ctx, req.GetEmail())
// 	if err != nil {
// 		return nil, status.Error(codes.NotFound, "User does not exist")
// 	}

// 	err = helpers.CheckPassword(req.GetPassword(), user.Password)
// 	if err != nil {
// 		return nil, status.Error(codes.NotFound, "User does not exist")
// 	}

// 	generateToken, err := s.repo.GenerateToken(ctx, user)
// 	if err != nil {
// 		return nil, status.Error(codes.Internal, "Internal")
// 	}

// 	return &userpb.LoginResponse{
// 		UserId:       user.ID,
// 		AccessToken:  generateToken.AccessToken,
// 		RefreshToken: generateToken.RefreshToken,
// 	}, nil
// }

// func (s *UserService) RefreshToken(ctx context.Context, req *userpb.RefreshTokenRequest) (*userpb.RefreshTokenResponse, error) {

// 	tok, err := s.repo.CheckToken(ctx, req.RefreshToken)
// 	if err != nil {
// 		return nil, status.Error(codes.NotFound, "User does not exist")
// 	}

// 	return &userpb.RefreshTokenResponse{
// 		AccessToken:  tok.AccessToken,
// 		RefreshToken: tok.RefreshToken,
// 	}, nil
// }

// func (h *UserService) RecoveryUser(ctx context.Context, req *userpb.RecoveryUserRequest) (*userpb.PayloadWithUserID, error) {

// 	user, err := h.repo.GetUserByEmail(ctx, req.GetUsername())
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = h.repo.RequestPasswordReset(ctx, user.UserID, user.Email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &userpb.PayloadWithUserID{
// 		UserId: user.UserID,
// 	}, nil
// }

// func (h *UserService) PasswordReset(ctx context.Context, req *userpb.PasswordResetRequest) (*emptypb.Empty, error) {

// 	dto := &usermd.PasswordResetUserDTO{
// 		UserID:   req.UserId,
// 		OTP:      req.Otp,
// 		Password: req.Password,
// 	}
// 	err := h.repo.PasswordReset(ctx, dto)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &emptypb.Empty{}, nil
// }
