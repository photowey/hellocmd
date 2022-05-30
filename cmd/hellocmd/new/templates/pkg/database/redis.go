package database

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	perrors "github.com/pkg/errors"

	rediz "codeup.aliyun.com/uphicoo/gokit/pkg/database/redis"

	"uphicoo.com/uphicoo/project-template/internal/config"
)

var _rt *rediz.RedizTemplate

// RedisClientInit 初始化 RedizTemplate 实例
func RedisClientInit(conf config.RedisConfig) (err error) {
	// 不实例化
	if !conf.Enabled {
		return nil
	}

	_rt, err = rediz.NewRedizTemplate(populateRedisAddress(conf), conf.Password, redis.DialDatabase(conf.DB))
	err = _rt.Ping()
	if err != nil {
		return perrors.Errorf("database.redis 数据库 PING 失败,请核实配置信息:%w", err)
	}

	return err
}

func RedisTemplate() *rediz.RedizTemplate {
	return _rt
}

func populateRedisAddress(conf config.RedisConfig) string {
	return fmt.Sprintf("%s:%d", conf.Host, conf.Port)
}
