package opt

import (
	"github.com/go-redis/redis"
)

//NewRedisClient  redis连接
func NewRedisClient() (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
