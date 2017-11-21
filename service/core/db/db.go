package db

import "gopkg.in/mgo.v2"
import "github.com/garyburd/redigo/redis"

const (
	url = "127.0.0.1:9500"
)

func MongoDB() (*mgo.Session, *mgo.Database, error) {
	session, err := mgo.Dial(url)
	if err == nil {
		db := session.DB("zhihu")
		return session, db, nil
	}

	return nil, nil, err
}

func RedisConnect() (redis.Conn, error) {
	return redis.Dial("tcp", "127.0.0.1:6379")
}
