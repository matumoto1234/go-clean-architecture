package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/matumoto1234/go-clean-architecture/app/model"
	"github.com/matumoto1234/go-clean-architecture/app/usecase"
)

// Direction of dependence : Repository -> Service
type UserRepository interface {
	FindUserByID(id string) (model.User, error)
	InsertUserByName(id, name string) error
}

// Direction of dependence : Service -> UseCase
type userServiceImpl struct {
	ur UserRepository
}

func (us userServiceImpl) FindUserByID(id string) (model.User, error) {
	return us.ur.FindUserByID(id)
}

func (us userServiceImpl) CreateUserByName(name string) error {
	// user id := version 4 UUID
	uuid, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("UserService.CreateUserByName: %w", err)
	}

	return us.ur.InsertUserByName(uuid.String(), name)
}

func NewUserService(ur UserRepository) usecase.UserService {
	return &userServiceImpl{
		ur: ur,
	}
}
