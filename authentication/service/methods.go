package service

import (
	"context"

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
		Id:        signUpUser.ID,
		Email:     signUpUser.Email,
		Uuid:      signUpUser.Uuid,
		Token:     signUpUser.Token,
		CreatedAt: signUpUser.CreatedAt,
	}, nil
}

func (s *authService) SignIn(ctx context.Context, req *proto.SignInRequest) (*proto.SignInResponse, error) {
	return &proto.SignInResponse{}, nil
}
