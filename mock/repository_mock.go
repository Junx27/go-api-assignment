package mock

import (
	"errors"
	"go-api-assignment/model"
)

type MockUserRepository struct {
	Users map[int]model.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{Users: make(map[int]model.User)}
}

func (m *MockUserRepository) AddUser(user model.User) error {
	m.Users[user.ID] = user
	return nil
}

func (m *MockUserRepository) GetUserByID(id int) (model.User, error) {
	user, exists := m.Users[id]
	if !exists {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}

func (m *MockUserRepository) GetUserByName(name string) (model.User, error) {
	for _, user := range m.Users {
		if user.Name == name {
			return user, nil
		}
	}
	return model.User{}, nil
}
