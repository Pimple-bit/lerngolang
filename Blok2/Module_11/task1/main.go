package storage

import "sync"

type StorageRWMutex struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewStorageRWMutex() *StorageRWMutex {
	return &StorageRWMutex{
		data: make(map[string]string),
	}
}

func (s *StorageRWMutex) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
}

func (s *StorageRWMutex) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.data[key]
	return val, ok
}