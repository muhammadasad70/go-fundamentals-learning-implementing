package main

import (
	"sync"
)

type UserStore struct {
	mu     sync.RWMutex
	users  map[int]User
	nextID int
}

func NewUserStore() *UserStore {
	return &UserStore{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func (s *UserStore) CreateUser(name, email string) User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := User{
		ID:    s.nextID,
		Name:  name,
		Email: email,
	}
	s.users[user.ID] = user
	s.nextID++
	return user

}
