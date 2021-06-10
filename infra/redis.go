package infra

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	pool *redis.Pool
}

type RedisConn struct {
	conn redis.Conn
}

func ConnectRedis() *Redis {
	return &Redis{
		pool: &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
			Dial: func() (redis.Conn, error) { return redis.Dial("tcp", ":6379") },
		},
	}
}

func (r *Redis) Close() error {
	return r.pool.Close()
}

func (r *Redis) Conn() *RedisConn {
	return &RedisConn{
		conn: r.pool.Get(),
	}
}

func (r *RedisConn) GetStruct(key string, dest interface{}) error {
	reply, err := redis.Bytes(r.conn.Do("GET", key))
	if err != nil {
		return err
	}
	return json.Unmarshal(reply, dest)
}

func (r *RedisConn) SetStruct(key string, data interface{}) (bool, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return false, err
	}
	return redis.Bool(r.conn.Do("SET", key, bytes))
}

func (r *RedisConn) GetString(key string) (string, error) {
	return redis.String(r.conn.Do("GET", key))
}

func (r *RedisConn) SetString(key, value string) error {
	ok, err := redis.Bool(r.conn.Do("SET", key, value))
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("Redis set failed")
	}

	return nil
}

func (r *Redis) RunWithCache(key string, dest interface{}, procedure func(dest interface{}) error) error {
	conn := r.Conn()
	defer conn.conn.Close()

	if err := conn.GetStruct(key, dest); err == nil {
		return nil
	}

	if err := procedure(dest); err != nil {
		return err
	}

	if _, err := conn.SetStruct(key, dest); err != nil {
		return err
	}

	return nil
}
