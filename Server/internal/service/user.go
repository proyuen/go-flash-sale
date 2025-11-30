package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/proyuen/flashSale/Server/internal/db"
	"github.com/proyuen/flashSale/Server/internal/util"
)

func (s *service) CreateUser(ctx context.Context, req CreateUserRequest) (UserResponse, error) {
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return UserResponse{}, fmt.Errorf("password encryption failed: %w", err)
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}

	user, err := s.store.CreateUser(ctx, arg)
	if err != nil {
		return UserResponse{}, err
	}

	return newUserResponse(user), nil
}

func (s *service) LoginUser(ctx context.Context, req LoginUserRequest) (LoginUserResponse, error) {
	user, err := s.store.GetUser(ctx, req.Username)
	if err != nil {
		return LoginUserResponse{}, err
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		return LoginUserResponse{}, fmt.Errorf("password incorrect: %w", err)
	}

	accessToken, err := s.tokenMaker.CreateToken(user.Username, s.config.AccessTokenDuration)
	if err != nil {
		return LoginUserResponse{}, fmt.Errorf("failed to create access token: %w", err)
	}

	return LoginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}, nil
}

func (s *service) GetUser(ctx context.Context, username string) (UserResponse, error) {
	user, err := s.store.GetUser(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return UserResponse{}, fmt.Errorf("user not found: %w", err)
		}
		return UserResponse{}, err
	}
	return newUserResponse(user), nil
}

func newUserResponse(user db.User) UserResponse {
	return UserResponse{
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time,
	}
}
