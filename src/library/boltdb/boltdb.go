package boltdb

import (
	"github.com/boltdb/bolt"
	"time"
	"fp/util/log"
	"github.com/davecgh/go-spew/spew"
	"fmt"
)

func Get() {
	db, err := bolt.Open("my.db", 0755, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	k := []byte("MyKey")
	var v []byte
	db.Update(func(tx *bolt.Tx) error {
		b, err1 := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err1 != nil {
			return fmt.Errorf("create bucket :%s", err1)
		}

		// b.Put(k, []byte(time.Now().String()))
		v = b.Get(k)
		return nil
	})
	spew.Dump(string(v))
	defer db.Close()
}
