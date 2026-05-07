package auth

import (
	"refresh_token/config"
	"refresh_token/internal/user"
	"refresh_token/pkg/encrypt"
	"refresh_token/pkg/token"
	"time"
)

type service struct {
	repo       Repository
	userRepo   user.Repository
	accessMgr  token.AccessTokenManager
	refreshMgr token.RefreshTokenManager
}

func NewService(repo Repository, userRepo user.Repository, accessMgr token.AccessTokenManager, refreshMgr token.RefreshTokenManager) *service {
	return &service{
		repo:       repo,
		userRepo:   userRepo,
		accessMgr:  accessMgr,
		refreshMgr: refreshMgr,
	}
}

func (s *service) Login(input *LoginRequestDTO) (*LoginResponseDTO, error) {
	user, err := s.userRepo.GetByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	if err := encrypt.Verify(input.Password, user.Password); err != nil {
		return nil, ErrInvalidCredentials
	}

	accessToken, err := s.accessMgr.GenerateToken(token.Payload{
		UserID: user.ID,
	})
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.refreshMgr.GenerateSecret()
	if err != nil {
		return nil, err
	}

	refreshTokenModel := &RefreshToken{
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(config.RefreshTokenTTL),
	}

	if err := refreshTokenModel.SetSecret(refreshToken); err != nil {
		return nil, err
	}

	if err := s.repo.Create(refreshTokenModel); err != nil {
		return nil, err
	}

	return &LoginResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
