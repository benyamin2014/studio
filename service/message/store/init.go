// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/8.

package store

import (
	"duguying/studio/g"
	"github.com/boltdb/bolt"
	"log"
)

var (
	boltDB *bolt.DB
)

func InitBoltDB() {
	dbPath := g.Config.Get("boltdb", "path", "performance.db")
	// open db
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		boltDB = db
	}

	initBucket()
}

func initBucket() error {
	tx, err := boltDB.Begin(true)
	if err != nil {
		return err
	}

	_, err = tx.CreateBucketIfNotExists([]byte("performance"))
	if err != nil {
		return err
	}

	_, err = tx.CreateBucketIfNotExists([]byte("agent"))
	if err != nil {
		return err
	}

	return tx.Commit()
}

func put(bucket string, key string, value []byte) error {
	tx, err := boltDB.Begin(true)
	if err != nil {
		return err
	}

	bkt := tx.Bucket([]byte(bucket))

	err = bkt.Put([]byte(key), value)
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}