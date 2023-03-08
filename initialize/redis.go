package initialize

import (
	"app_ws/global"
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"os"
)

var conn *redis.Client

// Redis Init 初始化连接
func Redis() (conn *redis.Client) {
	redisCfg := global.GVA_CONFIG.Redis
	conn = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})

	_, err := conn.Ping().Result()
	if err != nil {
		fmt.Printf("redis connect ping failed, err:%v\n", err)
		global.GVA_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
		os.Exit(0)
		return nil
	}
	return conn
}
