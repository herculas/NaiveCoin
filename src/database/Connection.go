package database

import (
	"github.com/boltdb/bolt"
	"log"
	"os"
)

const dbName = "data/naivechain.db"

func establishConnection() *bolt.DB {
	database, err := bolt.Open(dbName, 0600, nil)
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
	_, err := os.Stat(dbName)
	return !os.IsNotExist(err)
}

func Drop() {
	if err := os.Remove(dbName); err != nil {
		log.Panic(err)
	}
}