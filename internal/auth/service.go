package auth

import (
	"refresh_token/pkg/token"
)

type service struct {
	repo       Repository
	accessMgr  token.AccessTokenManager
	refreshMgr token.RefreshTokenManager
}

func NewService(repo Repository, accessMgr token.AccessTokenManager, refreshMgr token.RefreshTokenManager) *service {
	return &service{
		repo:       repo,
		accessMgr:  accessMgr,
		refreshMgr: refreshMgr,
	}
}
