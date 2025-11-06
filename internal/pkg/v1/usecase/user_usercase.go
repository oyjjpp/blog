package usecase

import (
	"api/internal/pkg/v1/domain"
)

func NewUserUseCase(user domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		user: user,
	}
}

type userUseCase struct {
	user domain.UserRepository
}
