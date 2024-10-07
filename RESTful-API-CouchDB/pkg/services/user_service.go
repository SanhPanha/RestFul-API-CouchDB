package services

import (
	"RESTful-API/domain"
	"RESTful-API/pkg/repositories"
)

type User = domain.User

func CreateUser(user User) error {
    // Business logic before saving
    return repositories.CreateUser(user)
}

// Other service methods...
