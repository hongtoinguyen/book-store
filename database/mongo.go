package database

import (
	"book-store/utils"
	"time"

	"gopkg.in/mgo.v2"
)

var db *mgo.Database

const (
	COL_CUSTOMERS = "customers"
)

//InitMongo ...
func InitMongo() {

	info := &mgo.DialInfo{
		Addrs:    []string{utils.GetEnv("MONGO_HOST", "localhost:27017")},
		Timeout:  60 * time.Second,
		Database: utils.GetEnv("MONGO_DATABASE", "book_store"),
		Username: utils.GetEnv("MONGO_USERNAME", "admin"),
		Password: utils.GetEnv("MONGO_PASSWORD", "bs_1234"),
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	session.Ping()

	db = session.DB(utils.GetEnv("MONGO_DATABASE", "book_store"))
}

//GetMongoDB ...
func GetMongoDB() *mgo.Database {
	return db
}
