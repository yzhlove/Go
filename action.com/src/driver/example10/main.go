package driver

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

const (
	fishingRankExpireTime = 7 * 24 * 60 * 60

	redisAddr     = "127.0.0.1"
	redisPort     = "6379"
	redisPassword = ""
	redisDBIndex  = 0
	redisPoolSize = 20
)

//RedisOpt redis客户端结构
type RedisOpt struct {
	Client *redis.Client
	Mutex  sync.Mutex
}

//RedisQueryModel 分页模型
type RedisQueryModel struct {
	Key   string
	Page  int64
	Total int64
}

//RedisQueryResult 查询模型
type RedisQueryResult struct {
	Page         int64
	Total        int64
	TotalEntries int64
	TotalPage    int64
	Result       []redis.Z
}

//GetRedisConnect redis连接
type GetRedisConnect func() (*RedisOpt, error)

//NewRedisClient 创建一个新的客户端
func NewRedisClient() GetRedisConnect {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr + ":" + redisPort,
		Password: redisPassword,
		DB:       redisDBIndex,
		PoolSize: redisPoolSize,
	})
	fmt.Printf("What are you doing ?")
	_, err := client.Ping().Result()
	return func() (*RedisOpt, error) {
		if err != nil {
			return nil, err
		}
		return &RedisOpt{Client: client}, nil
	}
}

func main() {

	redisConnect := NewRedisClient()
	if _, err := redisConnect(); err != nil {
		log.Fatalln(err)
		return
	}

}

//FindPage 分页
func (opt *RedisOpt) FindPage(model *RedisQueryModel) (*RedisQueryResult, bool) {
	if model.Page == 0 {
		model.Page = 1
	}
	if model.Total == 0 {
		model.Total = 20
	}
	totalEntries, err := opt.Client.ZCard(model.Key).Result()
	if err != nil {
		log.Fatalln(err)
		return nil, false
	}
	totalPage := int(totalEntries / model.Total)
	if totalEntries%model.Total == 0 {
		totalPage++
	}
	start := (model.Page - 1) * model.Total
	end := model.Page*model.Total - 1
	if results, err := opt.Client.ZRangeWithScores(model.Key, start, end).Result(); err == nil {
		return &RedisQueryResult{Page: model.Page, Total: int64(len(results)),
			TotalPage: int64(totalPage), TotalEntries: totalEntries, Result: results}, true
	} else {
		log.Fatalln(err)
	}
	return nil, false
}

//GetFishing 获取数据
func (opt *RedisOpt) GetFishing(model *RedisQueryModel) (map[string]float64, bool) {
	if result, ok := opt.FindPage(model); ok {
		dataMap := make(map[string]float64)
		for _, v := range result.Result {
			dataMap[v.Member.(string)] = v.Score
		}
		return dataMap, true
	}
	return nil, false
}

//InsertFishing 插入数据
func (opt *RedisOpt) InsertFishing(key, userUID string, value float64) (float64, bool) {

	opt.Mutex.Lock()
	defer opt.Mutex.Unlock()

	if result, err := opt.Client.ZIncrBy(key, value, userUID).Result(); err == nil {
		if ok, err := opt.Client.Expire(key, fishingRankExpireTime).Result(); ok {
			return result, true
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(err)
	}
	return 0, false
}

//GetFishingKey 获取key
func GetFishingKey(platform, style string) (string, bool) {
	platform = strings.TrimSpace(platform)
	style = strings.TrimSpace(style)
	if platform == "" || style == "" {
		return "", false
	}
	title := "finish:day_rank"
	timeFormat := time.Now().Format("20060102")
	keys := []string{title, platform, style, timeFormat}
	return strings.Join(keys, "_"), true
}
