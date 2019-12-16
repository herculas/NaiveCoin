package database

import (
	"github.com/boltdb/bolt"
	"log"
)

func Update(key []byte, value []byte) {
	var database = establishConnection()
	defer database.Close()
	err := database.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			log.Panic(err)
		}
		if bucket != nil {
			if err := bucket.Put(key, value); err != nil {
				log.Panic(err)
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func Retrieve(key []byte) []byte {
	var res []byte
	var database = establishConnection()
	defer database.Close()
	err := database.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket != nil {
			var content = bucket.Get(key)
			res = make([]byte, len(content), len(content))
			copy(res, content)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return res
}

func Delete(key []byte) {
	var database = establishConnection()
	defer database.Close()
	err := database.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket != nil {
			if err := bucket.Delete(key); err != nil {
				log.Panic(err)
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func DeleteBucket() {
	var database = establishConnection()
	defer database.Close()
	err := database.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket != nil {
			if err := tx.DeleteBucket([]byte(bucketName)); err != nil {
				log.Panic(err)
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}