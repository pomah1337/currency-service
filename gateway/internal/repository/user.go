package repository

import (
	"currencyService/gateway/internal/dto"
	"errors"
	"sync"
)

type UserRepository struct {
	users map[string]string
	mu    *sync.RWMutex
}

func NewUserRepository() UserRepository {
	return UserRepository{
		users: make(map[string]string),
		mu:    &sync.RWMutex{},
	}
}

func (r *UserRepository) SaveUser(login, password string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, contains := r.users[login]; contains {
		return errors.New("username \"" + login + "\" is occupied")
	}
	r.users[login] = password

	return nil
}

func (r *UserRepository) FindUser(login string) (dto.UserEntity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	password, contains := r.users[login]
	if !contains {
		return dto.UserEntity{}, errors.New("user \"" + login + "\" not found")
	}

	return dto.UserEntity{Login: login, Password: password}, nil
}
