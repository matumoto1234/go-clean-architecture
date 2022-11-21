package repository

import (
	"github.com/matumoto1234/go-clean-architecture/app/model"
	"github.com/matumoto1234/go-clean-architecture/app/service"
)

// Direction of dependence : Repository -> Service
type userRepositoryImpl struct {
	// This struct has db connection or something
	// e.g. rdb *rdb.Client
}

func (ur userRepositoryImpl) FindUserByID(id string) (model.User, error) {
	// Use the ORM or something to get the user from the database
	// The models or db will be defined in app/db/rdb/models/user.go

	return model.NewUser(
		id,
		"name",
	), nil
}

func (ur userRepositoryImpl) InsertUserByName(id, name string) error {
	// Use the ORM or something to insert the user to the database
	// The models or db will be defined in app/db/rdb/models/user.go

	return nil
}

func NewUserRepository() service.UserRepository {
	return &userRepositoryImpl{}
}
