package service

import (
	"context"
	"time"

	"github.com/aalsa16/golang-microservices/authentication/types"
	"github.com/aalsa16/golang-microservices/proto"
	"github.com/aalsa16/golang-microservices/utils"
)

func (s *authService) SignUp(ctx context.Context, req *proto.SignUpRequest) (*proto.User, error) {
	encryptedPassword, err := utils.EncryptPassword(req.Password)

	if err != nil {
		return &proto.User{}, err
	}

	user := &types.SignUpRequest{
		Email:    req.Email,
		Password: encryptedPassword,
	}

	signUpUser, err := s.database.SaveUser(user)

	if err != nil {
		return nil, err
	}

	return &proto.User{
		Id:    signUpUser.ID,
		Email: signUpUser.Email,
		Uuid:  signUpUser.Uuid,
	}, nil
}

func (s *authService) SignIn(ctx context.Context, req *proto.SignInRequest) (*proto.SignInResponse, error) {
	signInUser, err := s.database.GetUser(req.Email)

	if err != nil {
		return nil, err
	}

	err = utils.VerifyPassword(signInUser.Password, req.Password)

	if err != nil {
		return nil, err
	}

	accessToken, err := utils.NewToken(signInUser.Uuid, time.Now().Add(time.Minute*15))

	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.NewToken(signInUser.Uuid, time.Now().Add(time.Hour*24*7))

	if err != nil {
		return nil, err
	}

	err = s.database.SaveRefreshToken(refreshToken, signInUser.Uuid)

	if err != nil {
		return nil, err
	}

	user := &proto.User{
		Id:    signInUser.ID,
		Email: signInUser.Email,
		Uuid:  signInUser.Uuid,
	}

	return &proto.SignInResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) RefreshToken(ctx context.Context, req *proto.RefreshTokenRequest) (*proto.RefreshTokenResponse, error) {
	token, err := utils.RefreshToken(req.RefreshToken)

	if err != nil {
		return nil, err
	}

	return &proto.RefreshTokenResponse{
		AccessToken: token,
	}, nil
}
