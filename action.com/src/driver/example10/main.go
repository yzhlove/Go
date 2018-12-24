package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

const (
	fishingRankExpireTime = time.Duration(7*24*60*60) * time.Second

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

//NewRedisClient 创建一个新的客户端
func NewRedisClient() (*RedisOpt, bool) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr + ":" + redisPort,
		Password: redisPassword,
		DB:       redisDBIndex,
		PoolSize: redisPoolSize,
	})
	if _, err := client.Ping().Result(); err != nil {
		return nil, false
	}
	return &RedisOpt{Client: client}, true
}

func main() {

	opt, ok := NewRedisClient()
	if !ok {
		fmt.Printf("Create Redis Client Error!")
		return
	}
	key := GetFishingKey("yuewan", "gold")
	fmt.Printf("key = %v \n", key)
	for k, v := range map[string]float64{"123": 123, "456": 456, "789": 789} {
		go opt.InsertFishing(key, k, v)
	}

	time.Sleep(1 * time.Second)

	dataMap, ok := opt.GetFishing(&RedisQueryModel{Key: key, Page: 1, Total: 20})
	if !ok {
		fmt.Printf("Get Data Error!")
		return
	}
	fmt.Printf("dataMap = %v \n", dataMap)

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
	if results, err := opt.Client.ZRangeWithScores(model.Key, start, end).Result(); err != nil {
		log.Fatalln(err)
	} else {
		return &RedisQueryResult{Page: model.Page, Total: int64(len(results)),
			TotalPage: int64(totalPage), TotalEntries: totalEntries, Result: results}, true
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

	if result, err := opt.Client.ZIncrBy(key, value, userUID).Result(); err != nil {
		log.Fatalln(err)
	} else {
		if ok, err := opt.Client.Expire(key, fishingRankExpireTime).Result(); !ok {
			log.Fatalln(err)
		} else {
			return result, true
		}
	}
	return 0, false
}

//GetFishingKey 获取key
func GetFishingKey(platform, style string) string {
	platform = strings.TrimSpace(platform)
	style = strings.TrimSpace(style)
	if platform == "" || style == "" {
		return ""
	}
	title := "finish:day_rank"
	timeFormat := time.Now().Format("20060102")
	keys := []string{title, platform, style, timeFormat}
	return strings.Join(keys, "_")
}
