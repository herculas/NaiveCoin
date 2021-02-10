package database

import (
	"github.com/boltdb/bolt"
	"log"
	"naivechain/src/config"
	"naivechain/src/util/validator"
	"os"
)

func establishConn() *bolt.DB {
	if err := validator.IsDirExist(config.DatasourceDir); err != nil {
		log.Panic(err)
	}
	database, err := bolt.Open(config.DatabaseURI, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return database
}

func closeConn(db *bolt.DB) {
	if err := db.Close(); err != nil {
		log.Panic(err)
	}
}

func dropDatabase() {
	if err := os.Remove(config.DatabaseURI); err != nil {
		log.Panic(err)
	}
}