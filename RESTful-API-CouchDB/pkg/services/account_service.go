package services

import (
	"RESTful-API/domain"
	"RESTful-API/pkg/repositories"
)

type Account = domain.Account

func CreateAccount(account Account) error {
    // Business logic before saving
    return repositories.CreateAccount(account)
}

// Other service methods...
