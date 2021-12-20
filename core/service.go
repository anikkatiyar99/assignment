package core

import (
	mysvc "github.com/anikkatiyar99/mysvc"
)

type service struct {
	m map[int64]mysvc.User
}

// NewService instantiates a new Service.
func NewService() mysvc.Service {
	return &service{
		m: map[int64]mysvc.User{
			1: {
				ID:      1,
				Fname:   "Rohit",
				City:    "Bangalore",
				Phone:   1234567890,
				Height:  5.8,
				Married: true,
			},
			2: {
				ID:      2,
				Fname:   "Ajay",
				City:    "Mumbai",
				Phone:   987654321,
				Height:  6.1,
				Married: false,
			},
			3: {
				ID:      3,
				Fname:   "Karan",
				City:    "Delhi",
				Phone:   987654321,
				Height:  6.1,
				Married: false,
			},
		},
	}
}

// GetUser returns a user by ID.
func (s *service) GetUser(id int64) (result mysvc.User, err error) {
	if result, ok := s.m[id]; ok {
		return result, nil
	}

	return result, mysvc.ErrNotFound
}

// GetUsers returns a list of all users.
func (s *service) GetUsers(ids []int64) (result map[int64]mysvc.User, err error) {
	result = map[int64]mysvc.User{}

	for _, id := range ids {
		if u, ok := s.m[id]; ok {
			result[id] = u
		}
	}

	return
}
