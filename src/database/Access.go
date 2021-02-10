package database

import (
	"github.com/boltdb/bolt"
	"log"
)

func Update(bucketName string, key []byte, value []byte) {
	var db = establishConn()
	defer closeConn(db)

	err := db.Update(func(tx *bolt.Tx) error {
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

func Retrieve(bucketName string, key []byte) []byte {
	var res []byte
	var db = establishConn()
	defer closeConn(db)
	err := db.View(func(tx *bolt.Tx) error {
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

func Iterate(bucketName string, exec func(k, v []byte) error) {
	db := establishConn()
	defer closeConn(db)
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket != nil {
			if err := bucket.ForEach(exec); err != nil {
				log.Panic(err)
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func Delete(bucketName string, key []byte) {
	var db = establishConn()
	defer closeConn(db)
	err := db.Update(func(tx *bolt.Tx) error {
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

func CheckBucket(bucketName string) bool {
	var state = true
	db := establishConn()
	defer closeConn(db)
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			state = false
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return state
}

func DropBucket(bucketName string) {
	db := establishConn()
	defer closeConn(db)
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket != nil {
			return tx.DeleteBucket([]byte(bucketName))
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}