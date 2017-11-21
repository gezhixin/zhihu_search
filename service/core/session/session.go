package session

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/json-iterator/go"
	"gopkg.in/mgo.v2/bson"
	"xst_Cloud/service/core/db"
)

type Session struct {
	Id     string
	UserId string
	DevId  string
	Data   map[string]interface{}
}

func NewSession() (*Session, error) {
	redisCon, err := db.RedisConnect()
	if err == nil {
		defer redisCon.Close()
		var session Session
		session.Id = bson.NewObjectId().Hex()
		session.UserId = ""
		session.DevId = ""
		session.Data = make(map[string]interface{})
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		sessionData, _ := json.Marshal(&session)
		_, err := redisCon.Do("SET", session.Id, sessionData)
		if err != nil {
			return nil, err
		}
		redisCon.Do("EXPIRE", session.Id, 60*60*24*30)

		return &session, nil
	} else {
		return nil, err
	}
}

func GetSession(sessionId string) (*Session, error) {
	redisCon, erro := db.RedisConnect()
	if erro == nil {
		defer redisCon.Close()
		value, err := redis.Bytes(redisCon.Do("GET", sessionId))
		if err != nil {
			return nil, err
		}
		var session Session
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		json.Unmarshal(value, &session)
		return &session, nil
	} else {
		return nil, erro
	}
}

func GetSessionByUserId(userId string) (*Session, error) {
	redisCon, erro := db.RedisConnect()
	if erro == nil {
		defer redisCon.Close()
		value, err := redis.String(redisCon.Do("GET", userId))
		if err != nil {
			return nil, err
		}
		return GetSession(value)
	} else {
		return nil, erro
	}
}

func DeleteSession(sessionId string) {
	redisCon, erro := db.RedisConnect()
	if erro == nil {
		defer redisCon.Close()
		redisCon.Do("DEL", sessionId)
	}
}

func (s *Session) SetExpireTime(time int64) {
	redisCon, erro := db.RedisConnect()
	if erro == nil {
		defer redisCon.Close()
		redisCon.Do("EXPIRE", s.Id, time)
	}
}

func (s *Session) Delete() {
	redisCon, erro := db.RedisConnect()
	if erro == nil {
		defer redisCon.Close()
		redisCon.Do("DEL", s.DevId)
		redisCon.Do("DEL", s.UserId)
		redisCon.Do("DEL", s.Id)
	}
}

func (s *Session) Update() error {
	redisCon, err := db.RedisConnect()
	if err == nil {
		defer redisCon.Close()
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		sessionData, _ := json.Marshal(&s)
		_, err := redisCon.Do("SET", s.Id, sessionData)

		if len(s.UserId) > 0 {
			redisCon.Do("SET", s.UserId, s.Id)
			redisCon.Do("EXPIRE", s.UserId, 60*60*24*30)
		}

		if len(s.DevId) > 0 {
			redisCon.Do("SET", s.DevId, s.Id)
			redisCon.Do("EXPIRE", s.DevId, 60*60*24*30)
		}

		if err != nil {
			return err
		}

		return nil
	} else {
		return err
	}
}

func (s *Session) SetKeyValue(key string, value interface{}) error {

	redisCon, err := db.RedisConnect()
	if err == nil {
		defer redisCon.Close()
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		if s.Data == nil {
			s.Data = make(map[string]interface{})
		}
		s.Data[key] = value
		fmt.Println("s ---> ")
		fmt.Println(s)
		sessionData, _ := json.Marshal(&s)
		_, err := redisCon.Do("SET", s.Id, sessionData)
		if err != nil {
			return err
		}

		return nil
	} else {
		return err
	}
}

func (s *Session) GetValue(key string) (interface{}, bool) {
	if v, ok := s.Data[key]; ok {
		return v, true
	} else {
		return "", false
	}
}
