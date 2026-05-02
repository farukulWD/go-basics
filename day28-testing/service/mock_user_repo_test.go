package service

import (
	"errors"

	"go-basics/day28-testing/domain"
)

// mockUserRepo is an in-memory fake of domain.UserRepository.
// Used only in tests — compiled only when running go test.
type mockUserRepo struct {
	users  []domain.User
	nextID uint
}

func newMockUserRepo() *mockUserRepo {
	return &mockUserRepo{nextID: 1}
}

func (m *mockUserRepo) Create(user *domain.User) error {
	for _, u := range m.users {
		if u.Email == user.Email {
			return errors.New("email already exists")
		}
	}
	user.ID = m.nextID
	m.nextID++
	m.users = append(m.users, *user)
	return nil
}

func (m *mockUserRepo) FindByEmail(email string) (*domain.User, error) {
	for _, u := range m.users {
		if u.Email == email {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m *mockUserRepo) FindByID(id uint) (*domain.User, error) {
	for _, u := range m.users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m *mockUserRepo) Update(user *domain.User) error {
	for i, u := range m.users {
		if u.ID == user.ID {
			m.users[i] = *user
			return nil
		}
	}
	return errors.New("user not found")
}

func (m *mockUserRepo) FindAll() ([]domain.User, error) {
	return m.users, nil
}
