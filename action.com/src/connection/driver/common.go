package driver

import (
	"database/sql"
	"fmt"

	"github.com/go-redis/redis"

	//导入sql包
	_ "github.com/lib/pq"
)

//RedisConn redis
type RedisConn struct {
	Addr     string
	PassWord string
	DB       int
}

//PostgresConn postgre
type PostgresConn struct {
	Host     string
	Port     int
	User     string
	PassWord string
	Dbname   string
}

//SsdbConn ssdb
type SsdbConn struct {
	Host string
	Port int
}

type conner interface {
	newClient() (interface{}, error)
}

//NewClient  创建一个客户端
func NewClient(it interface{}) interface{} {

	switch it.(type) {
	case *RedisConn:

		rd, err := it.(*RedisConn).newClient()
		if err != nil {
			panic(err)
		}
		return rd
	case *PostgresConn:
		pd, err := it.(*PostgresConn).newClient()
		if err != nil {
			panic(err)
		}
		return pd
	default:
		return nil
	}
}

func (rc *RedisConn) newClient() (interface{}, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     rc.Addr,
		Password: rc.PassWord,
		DB:       rc.DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (pc *PostgresConn) newClient() (interface{}, error) {

	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		pc.Host, pc.Port, pc.User, pc.Dbname, pc.PassWord)
	postgres, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	err = postgres.Ping()
	if err != nil {
		return nil, err
	}
	return postgres, nil
}
