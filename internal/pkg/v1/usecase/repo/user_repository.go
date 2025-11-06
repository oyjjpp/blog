package repo

import (
	"api/internal/pkg/v1/domain"
)

func NewUserRepository() domain.UserRepository {
	return &userRepository{}
}

type userRepository struct {
}
