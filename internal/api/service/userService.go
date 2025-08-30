package service

import (
	"gifma-backend/config"
	"gifma-backend/internal/helper"
	"gifma-backend/token"
	"gifma-backend/util"
	"time"

	"gifma-backend/internal/api/domain"
)

type UserService struct {
	// Repo       repository.UserRepository
	Config     config.AppConfig
	Crypto     helper.Crypto
	TokenMaker token.Maker
}

func (s *UserService) CreateUser(email string, hashedPassword string, role string) (*domain.User, error) {

	// TODO hash password
	hashedPassword, err := util.HashPassword(hashedPassword)
	if err != nil {
		return nil, err
	}

	// TODO register user

	registerdUser := &domain.User{
		Email:    email,
		Password: hashedPassword,
		Role:     role,
	}

	return registerdUser, nil
}

func (s *UserService) Login(email string, password string) (*domain.LoginUser, error) {
	// TODO get user by mail
	user := &domain.User{
		Email:    email,
		Password: password,
		Role:     "role",
	}

	// sample expire
	accessTokenExpiresAt := time.Now().Add(time.Hour)
	refreshTokenExpiresAt := time.Now().Add(time.Hour * 24)

	// TODO check password

	// TODO verify check

	// TODO generate access token

	// TODO generate refresh token

	// TODO check session

	// TODO get user information (cart, order, etc)

	result := &domain.LoginUser{
		User:                  user,
		AccessToken:           "accessToken",
		RefreshToken:          "refreshToken",
		AccessTokenExpiresAt:  accessTokenExpiresAt,
		RefreshTokenExpiresAt: refreshTokenExpiresAt,
	}

	return result, nil
}
