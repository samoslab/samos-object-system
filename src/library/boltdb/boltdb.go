package boltdb

import (
	"github.com/boltdb/bolt"
	"time"
	"log"
	"errors"
)

func List(path, bucket string) (map[string][]byte, error) {
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("open data file failed")
	}
	var result = make(map[string][]byte)
	err3 := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			err2 := errors.New("open bucket failed")
			log.Fatal(err2)
			return err2
		}
		b.ForEach(func(k, v []byte) error {
			result[string(k)] = v
			return nil
		})
		return nil
	})
	if err3 != nil{
		return nil, err3
	}
	var data = make(map[string][]byte)
	for k, v := range result {
		data[k] = []byte(string(v))
	}
	// 防止错误的内存
	return data, nil
}

// 读取内容
func Get(path, bucket, key string) ([]byte, error) {
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("open data file failed")
	}
	var v []byte
	err3 := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			err2 := errors.New("open bucket faild")
			log.Fatal(err2)
			return err2
		}
		v = b.Get([]byte(key))
		return nil
	})
	if err3 != nil{
		return nil, err3
	}
	// 防止错误的内存
	return []byte(string(v)), nil
}

// 设置内容
func Set(path, bucket, key string, value []byte) error {
	db, err := bolt.Open(path, 0755, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return errors.New("open data file failed")
	}
	result := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			log.Fatal(err)
			return errors.New("create or read bucket failed")
		}
		err3 := b.Put([]byte(key), value)
		if err3 != nil {
			log.Fatal(err3)
			return errors.New("put value failed")
		}
		return nil
	})
	return result
}
