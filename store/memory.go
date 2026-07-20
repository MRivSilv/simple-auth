package store

import (
	"sync"
)

type UserStore struct {
	mu sync.RWMutex
	users map[string]string
}

func NewUserStore() *UserStore{
	return &UserStore{users: make(map[string]string)}
}

func (s *UserStore) Create(username, hash string) bool{
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists :=s.users[username]; exists{
		return false
	}
	s.users[username] = hash
	return true
}

func (s *UserStore) Get(username string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	hash, ok  := s.users[username]
	return hash, ok
}

func (s *UserStore) Delete(username string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.users[username]; !exists {
		return false
	}
	delete(s.users, username)
	return true
}