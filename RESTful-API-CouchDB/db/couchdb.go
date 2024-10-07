package db

import (
    kivik "github.com/go-kivik/kivik/v4"
    _ "github.com/go-kivik/couchdb/v4"
    "log"
)

var CouchClient *kivik.Client

func InitCouchDB() {
    client, err := kivik.New("couch", "http://localhost:5984/")
    if err != nil {
        log.Fatal("Failed to connect to CouchDB:", err)
    }

    CouchClient = client
}