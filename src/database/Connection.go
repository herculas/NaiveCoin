package database

import (
	"github.com/boltdb/bolt"
	"log"
)

const databaseName = "src/data/naivechain.db"
const bucketName = "blocks"

func establishConnection() *bolt.DB {
	database, err := bolt.Open(databaseName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return database
}
