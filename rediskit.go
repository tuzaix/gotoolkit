package gotoolkit


import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type RedisPool struct {
	*redis.Pool
}

/**
	初始化redis池
 */
func NewRedisPool(host string, port, maxIdle, idleTimeout, connectTimeout, readTimeout, writeTimeout int) *RedisPool {
	_url := fmt.Sprintf("%s:%d", host, port)
	_idleTO := time.Duration(idleTimeout)*time.Millisecond
	_connectTO := time.Duration(connectTimeout)*time.Millisecond
	_readTO := time.Duration(readTimeout)*time.Millisecond
	_writeTO := time.Duration(writeTimeout)*time.Millisecond

	return &RedisPool{
		&redis.Pool{
			MaxIdle:     maxIdle,
			MaxActive: 	0,
			IdleTimeout: _idleTO,
			Dial: func() (redis.Conn, error) {
				var (
					conn redis.Conn
					err  error
				)
				if conn, err = redis.Dial("tcp", _url, redis.DialConnectTimeout(_connectTO), redis.DialReadTimeout(_readTO), redis.DialWriteTimeout(_writeTO)); err != nil {
					return nil, err
				}
				if _, err = conn.Do("EXISTS", rand.Int63n(1000)); err != nil {
					conn.Close()
					return nil, err
				}
				return conn, err
			},
			TestOnBorrow: func(conn redis.Conn, t time.Time) error {
				t = t.Add(_idleTO)
				if time.Now().Before(t) {
					return nil
				}
				var err error
				_, err = conn.Do("EXISTS", rand.Int63n(1000))
				return err
			},
			Wait: true,
		}}
}

func (this *RedisPool) Expire(args ...interface{}) (bool, error) {
	conn := this.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("EXPIRE", args...))
}

func (this *RedisPool) Bool(commandName string, args ...interface{}) (bool, error) {
	conn := this.Get()
	defer conn.Close()
	return redis.Bool(conn.Do(commandName, args...))
}

func (this *RedisPool) Bytes(commandName string, args ...interface{}) ([]byte, error) {
	conn := this.Get()
	defer conn.Close()
	return redis.Bytes(conn.Do(commandName, args...))
}

func (this *RedisPool) Float64(commandName string, args ...interface{}) (float64, error) {
	conn := this.Get()
	defer conn.Close()
	return redis.Float64(conn.Do(commandName, args...))
}

func (this *RedisPool) Int(commandName string, args ...interface{}) (int, error) {
	conn := this.Get()
	defer conn.Close()
	return redis.Int(conn.Do(commandName, args...))
}

func (this *RedisPool) Int64(commandName string, args ...interface{}) (int64, error) {
	conn := this.Get()
	defer conn.Close()
	return redis.Int64(conn.Do(commandName, args...))
}

func (this *RedisPool) Ints(commandName string, args ...interface{}) ([]int, error) {
	conn := this.Get()
	defer conn.Close()
	return redis.Ints(conn.Do(commandName, args...))
}

func (this *RedisPool) MultiBulk(commandName string, args ...interface{}) ([]interface{}, error) {
	conn := this.Get()
	defer conn.Close()
	return redis.Values(conn.Do(commandName, args...))
}

func (this *RedisPool) String(commandName string, args ...interface{}) (string, error) {
	conn := this.Get()
	defer conn.Close()
	return redis.String(conn.Do(commandName, args...))
}

func (this *RedisPool) Strings(commandName string, args ...interface{}) ([]string, error) {
	conn := this.Get()
	defer conn.Close()
	return redis.Strings(conn.Do(commandName, args...))
}

func (this *RedisPool) StringMap(commandName string, args ...interface{}) (map[string]string, error) {
	conn := this.Get()
	defer conn.Close()
	return redis.StringMap(conn.Do(commandName, args...))
}
