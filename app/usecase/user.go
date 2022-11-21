package usecase

import (
	"github.com/matumoto1234/go-clean-architecture/app/model"
)

// Direction of dependence : Service -> UseCase
type UserService interface {
	FindUserByID(id string) (model.User, error)
	CreateUserByName(name string) error
}

// Direction of dependence : Controller -> UseCase
type UserUseCase interface {
	GETUser(id string) (model.User, error)
	POSTUser(id string) error
}

type userUseCaseImpl struct {
	us UserService
}

func (uu userUseCaseImpl) GETUser(id string) (model.User, error) {
	user, err := uu.us.FindUserByID(id)
	if err != nil {
		return model.User{}, &Error{
			Err:  err,
			Kind: ErrInternalServerError,
		}
	}

	if user.IsEmpty() {
		return model.User{}, &Error{
			Err:  err,
			Kind: ErrNotFound,
		}
	}

	return user, nil
}

func (uu userUseCaseImpl) POSTUser(name string) error {
	err := uu.us.CreateUserByName(name)
	if err != nil {
		return &Error{
			Err: err,
			Kind: ErrInternalServerError,
		}
	}

	return nil
}

func NewUserUseCase(us UserService) UserUseCase {
	return &userUseCaseImpl{
		us: us,
	}
}
