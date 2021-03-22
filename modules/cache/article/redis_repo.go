package article

import (
	"encoding/json"
	"kumparan/business/article"
	"time"

	"github.com/go-redis/redis"
)

// RedisRepository ...
type RedisRepository struct {
	Redis *redis.Client
}

// NewCacheRepository ...
func NewCacheRepository(Redis *redis.Client) *RedisRepository {
	return &RedisRepository{
		Redis,
	}
}
func (client *RedisRepository) SetKey(key string, value interface{}, expired int64) bool {
	exp := time.Duration(expired) * time.Second
	err := client.Redis.Set(key, value, exp).Err()
	if err != nil {
		return false
	}

	return true
}

// Get ...
func (client *RedisRepository) Get(key string) string {
	val, err := client.Redis.Get(key).Result()
	if err != nil {
		return ""
	}
	return val
}

func (client *RedisRepository) MGet(keys []string) []article.Article {
	var ret []article.Article
	sc := client.Redis.MGet(keys...)
	ret, err := ToKeyVal(sc)
	if err != nil {
		return nil
	}

	return ret
}

func (client *RedisRepository) GetAllKeys() []string {
	var ret []string
	keys := client.Redis.Keys("*")
	keyRes, err := keys.Result()
	if err != nil {
		return nil
	}
	for _, key := range keyRes {
		ret = append(ret, key)
	}

	return ret
}

func ToKeyVal(cmd *redis.SliceCmd) ([]article.Article, error) {
	var artcls []article.Article
	var artcl article.Article
	xss, err := cmd.Result()
	if err != nil {
		return nil, err
	}

	for _, xs := range xss {
		json.Unmarshal([]byte(xs.(string)), &artcl)

		artcls = append(artcls, artcl)
	}

	return artcls, nil
}
