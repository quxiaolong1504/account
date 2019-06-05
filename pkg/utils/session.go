
/*

An in-redis session store implementation.

*/

package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/icza/session"
)

type redisStore struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) session.Store {
	return redisStore{
		client: client,
	}
}

func (r redisStore) Get (id string) session.Session {
	redisKey := r.GenRedisKey(id)
	data, err := r.client.Get(redisKey).Result()
	if err != nil {
		return nil
	}
	return r.DecodeSession([]byte(data))
}

func (r redisStore) Add (sess session.Session){
	redisKey := r.GenRedisKey(sess.ID())
	data := r.EncodeSession(sess)
	_, err := r.client.Set(redisKey, data, sess.Timeout()).Result()
	if err != nil {
		panic("")
	}
}

func (r redisStore) Remove(sess session.Session) {
	redisKey := r.GenRedisKey(sess.ID())
	r.client.Del(redisKey)
}

func (r redisStore) Close(){
	// do nothing here
}

func (r redisStore) GenRedisKey(id string) string {
	return fmt.Sprintf("auth:session:%s", id)
}

func (r redisStore) EncodeSession(sess session.Session) []byte {
	data, err := json.Marshal(sess)
	if err != nil {
		panic(err.Error())
	}
	return data
}

func (r redisStore) DecodeSession(data []byte) session.Session {

	sess := session.NewSession()
	err := json.Unmarshal(data, &sess)
	if err != nil {
		panic(err.Error())
	}
	return sess
}