package repositories

import (
	"RESTful-API/db"
	"RESTful-API/domain"
	"context"
)

func CreateAccount(account domain.Account) error {
	db := db.CouchClient.DB("accounts")

    _, err := db.Put(context.Background(), account.ID, account)
    return err
}

// Other repository methods...
