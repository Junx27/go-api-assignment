package repository

import (
	"fmt"
	"go-api-assignment/model"
)

type UserRepository interface {
	AddUser(user model.User) error
	GetUserByID(id int) (model.User, error)
	GetUserByName(name string) (model.User, error)
}

type UserRepositoryImpl struct {
	users map[int]model.User
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{
		users: make(map[int]model.User),
	}
}

func (repo *UserRepositoryImpl) AddUser(user model.User) error {
	repo.users[user.ID] = user
	return nil
}
func (repo *UserRepositoryImpl) GetUserByID(id int) (model.User, error) {
	user, exists := repo.users[id]
	if !exists {
		return model.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}
func (repo *UserRepositoryImpl) GetUserByName(name string) (model.User, error) {
	for _, user := range repo.users {
		if user.Name == name {
			return user, nil
		}
	}
	return model.User{}, nil
}
