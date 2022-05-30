package app

import (
	"codeup.aliyun.com/uphicoo/gokit/log"
	"codeup.aliyun.com/uphicoo/gokit/stringz"
	perrors "github.com/pkg/errors"

	"uphicoo.com/uphicoo/project-template/internal/config"
	"uphicoo.com/uphicoo/project-template/internal/config/configregistry"
	"uphicoo.com/uphicoo/project-template/pkg/database"
)

func Start(conf string) error {
	if stringz.IsBlankString(conf) {
		return perrors.Errorf("config file is empty")
	}

	// 1.加载: App 配置
	err := loadConfig(conf)
	if err != nil {
		return err
	}

	// 1.1.广播配置 Config 已经加载并初始化好了
	// TODO 基于事件的方式处理 - 直接 eventbus.Publish(ConfigFinishedEvent)
	configregistry.Broadcast()

	// 2.初始化: 系统日志 logger
	err = initLogger()
	if err != nil {
		return err
	}

	// 3.初始化: 数据库
	if err = database.RDBMSClientInit(config.DatabaseMap()); err != nil {
		return perrors.Errorf("app: init the rdbms error: %v", err)
	}

	// 4.初始化: Redis 缓存
	if err = database.RedisClientInit(config.Redis()); err != nil {
		return perrors.Errorf("app: init the redis error: %v", err)
	}

	// TODO

	return nil
}

// 加载: App 配置
func loadConfig(conf string) error {
	if err := config.LoadToml(conf); err != nil {
		return perrors.Errorf("app: load the config file:[%s] failed: %v", conf, err)
	}

	return nil
}

// 初始化: 系统日志 logger
func initLogger() error {
	if err := log.Init(config.Log()); err != nil {
		return perrors.Errorf("app: init app logger failed: %v", err)
	}
	return nil
}