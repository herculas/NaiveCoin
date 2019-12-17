package database

import (
	"github.com/boltdb/bolt"
	"log"
	"os"
)

const databaseName = "data/naivechain.db"
const bucketName = "blocks"

func establishConnection() *bolt.DB {
	database, err := bolt.Open(databaseName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return database
}

func Exists() bool {
	_, err := os.Stat(databaseName)
	return !os.IsNotExist(err)
}
