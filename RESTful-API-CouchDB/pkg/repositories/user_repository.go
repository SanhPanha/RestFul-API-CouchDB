package repositories

import (
	"RESTful-API/db"
	"RESTful-API/domain"
	"context"
)

func CreateUser(user domain.User) error {
    db := db.CouchClient.DB("users")

    _, err := db.Put(context.Background(), user.ID, user)
    return err
}

func GetUser(id string) (*domain.User, error) {
    // Fetch user by ID logic
    return nil, nil
}

// Other repository methods...
