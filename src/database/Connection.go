package database

import (
	"github.com/boltdb/bolt"
	"log"
	"os"
)

const DBName = "data/naivechain.db"
const BucketName = "blocks"

func establishConnection() *bolt.DB {
	database, err := bolt.Open(DBName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return database
}

func closeConnection(database *bolt.DB) {
	if err := database.Close(); err != nil {
		log.Panic(err)
	}
}

func IsExist() bool {
	_, err := os.Stat(DBName)
	return !os.IsNotExist(err)
}

func Drop() {
	if err := os.Remove(DBName); err != nil {
		log.Panic(err)
	}
}