package main

import "sync"

type Set interface {
	Set(s string)
	Get(s string) (string, bool)
	Delete(s string)
}

type set struct {
	mp map[string]struct{}
	mu sync.RWMutex
}

func NewSet() Set {
	return &set{
		mp: make(map[string]struct{}),
	}
}

func (s *set) Set(v string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.mp[v] = struct{}{}
}

func (s *set) Get(v string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.mp[v]
	return v, ok
}

func (s *set) Delete(v string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.mp, v)
}

func main() {

}
