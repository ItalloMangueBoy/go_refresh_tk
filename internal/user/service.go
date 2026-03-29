package user

import (
	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(input *CreateDTO) (*ResponseDTO, error) {
	user, err := input.ToModel()
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	resp := user.ToResponse()
	return &resp, nil
}

func (s *Service) GetByID(id uuid.UUID) (*ResponseDTO, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	resp := user.ToResponse()
	return &resp, nil
}
