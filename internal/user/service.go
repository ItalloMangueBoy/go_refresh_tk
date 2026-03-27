package user

import "refresh_token/internal/user/dto"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(input *dto.CreateInput) (*dto.UserResponse, error) {
	exists, err := s.repo.ExistsByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrUserAlreadyExists
	}

	user := &User{
		Name:  input.Name,
		Email: input.Email,
	}
	user.SetPassword(input.Password)

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	resp := user.ToResponse()
	return &resp, nil
}
