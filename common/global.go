package common

import (
	"github.com/auto-cdp/utils/etcd"
	"github.com/auto-cdp/utils/log"
	"github.com/auto-cdp/utils/redis"
	"time"
)

var (
	RedisConn    redis.RedisConn
	EtcdClient   *etcd.AfisServiceRegister
	AgentID      string
	ConfigKey    string
	EtcdEndpoint []string
	DialTimeout  time.Duration = 30
)

func InitRedis() {
	redispool := redis.NewRedisPool(Config().Redis.Host,
		Config().Redis.MaxIdle,
		Config().Redis.MaxActive,
		Config().Redis.Timeout)

	RedisConn = redis.NewRedisConn(redispool)
}

func InitEtcdClient() {
	asr := &etcd.AfisServiceRegister{}
	var err error
	EtcdClient, err = asr.NewAfisServiceRegister(EtcdEndpoint, DialTimeout)
	if err != nil {
		log.Slogger.Fatalf("etcd client init fail.[%s]", err)
	}
}
