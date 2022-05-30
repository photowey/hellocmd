package config

import (
	"fmt"
	"time"

	"codeup.aliyun.com/uphicoo/gokit/jsonz"
	"codeup.aliyun.com/uphicoo/gokit/log"
	"codeup.aliyun.com/uphicoo/gokit/pkg/loader"
	"codeup.aliyun.com/uphicoo/gokit/stringz"
	perrors "github.com/pkg/errors"
)

//
// 应用配置
//

type Profile = string

var (
	EnvLocal     = Profile("local")
	envTest      = Profile("test")
	envBenchmark = Profile("benchmark")
	envStage     = Profile("stage") // 预发布 final
	envProd      = Profile("prod")

	profiles = []Profile{EnvLocal, envTest, envBenchmark, envStage, envProd}
)

var (
	_conf   Config
	_inited bool
)

// Config App 配置对象
type Config struct {
	PrintConfig bool                `toml:"print_config" json:"printConfig" yaml:"printConfig"`      // 解析配置完成后是否打印配置信息
	Env         string              `toml:"env" json:"env" yaml:"env"`                               // 当前环境
	Host        string              `toml:"host" json:"host" yaml:"host"`                            // app 启动时监听的 host
	TimeOut     uint32              `toml:"start_check_health_time_out" json:"timOut" yaml:"timOut"` // app 启动时监听的 host
	Log         log.Config          `toml:"log" json:"log" yaml:"log"`                               // 日志配置
	DatabaseMap map[string]DBConfig `toml:"databases" json:"databases" yaml:"databases"`             // 数据库列表 配置 -> 操作多数据源场景
	RabbitMQ    RabbitMQConfig      `toml:"rabbitmq" json:"rabbitmq" yaml:"rabbitmq"`                // rabbitmq 配置
	Redis       RedisConfig         `toml:"redis" json:"redis" yaml:"redis"`                         // redis 配置
}

// durationWrapper 时间间隔包装
type durationWrapper struct {
	time.Duration
}

// DBConfig 数据库配置
type DBConfig struct {
	Enabled         bool            `toml:"enabled" json:"enabled" yaml:"enabled"`        // 是否开启数据库 false: 将不会实例化 数据库客户端实例
	Driver          string          `toml:"driver" json:"driver" yaml:"driver"`           // 数据库驱动
	Dbname          string          `toml:"dbname" json:"dbname" yaml:"dbname"`           // 连接的数据库名称
	Username        string          `toml:"username" json:"username" yaml:"username"`     // 数据库-用户名
	Password        string          `toml:"password" json:"password" yaml:"password"`     // 数据库-密码
	WriteHost       string          `toml:"write_host" json:"writeHost" yaml:"writeHost"` // 数据库-写请求-主机
	ReadHost        string          `toml:"read_host" json:"readHost" yaml:"readHost"`    // 数据库-读请求-主机
	WritePort       int             `toml:"write_port" json:"writePort" yaml:"writePort"` // 数据库-写请求--端口
	ReadPort        int             `toml:"read_port" json:"readPort" yaml:"readPort"`    // 数据库-读请求-端口
	ConnMaxIdleTime durationWrapper `toml:"conn_max_idle_time" json:"connMaxIdleTime" yaml:"connMaxIdleTime"`
	ConnMaxLifeTime durationWrapper `toml:"conn_max_life_time" json:"connMaxLifeTime" yaml:"connMaxLifeTime"`
	MaxIdleConns    int             `toml:"max_idle_conns" json:"maxIdleConns" yaml:"maxIdleConns"`
	MaxOpenConns    int             `toml:"max_open_conns" json:"maxOpenConns" yaml:"maxOpenConns"`
}

// RabbitMQConfig rabbitmq 配置
type RabbitMQConfig struct {
	Enabled      bool                         `toml:"enabled" json:"enabled" yaml:"enabled"` // 是否开启 rabbitmq 消息队列 false: 将不会实例化 rabbitmq 客户端实例
	Host         string                       `toml:"host" json:"host" yaml:"host"`
	Port         int                          `toml:"port" json:"port" yaml:"port"`
	User         string                       `toml:"user" json:"user" yaml:"user"`
	Password     string                       `toml:"password" json:"password" yaml:"password"`
	Vhosts       string                       `toml:"vhosts" json:"vhosts" yaml:"vhosts"`
	ExchangerMap map[string]RabbitMQExchanger `toml:"exchangers" json:"exchangers" yaml:"exchangers"`
}

type RabbitMQExchanger struct {
	Exchanger  string `toml:"exchanger" json:"exchanger" yaml:"exchanger"`
	RoutingKey string `toml:"routing_key" json:"routingKey" yaml:"routingKey"`
	Queue      string `toml:"queue" json:"queue" yaml:"queue"`
}

// RedisConfig Redis 配置
type RedisConfig struct {
	Enabled  bool   `toml:"enabled" json:"enabled" yaml:"enabled"` // 是否开启 Redis 缓存从 false: 将不会实例化 RedizTemplate实例
	Prefix   string `toml:"prefix" json:"prefix" yaml:"prefix"`
	Host     string `toml:"host" json:"host" yaml:"host"`
	Port     int64  `toml:"port" json:"port" yaml:"port"`
	Username string `toml:"username" json:"username" yaml:"username"`
	Password string `toml:"password" json:"password" yaml:"password"`
	DB       int    `toml:"db" json:"db" yaml:"db"`
}

func (d *durationWrapper) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))

	return err
}

// LoadToml 加载 toml 类型的配置
//
// 后续可能还会有 yaml 类型的配置 LoadYaml
//
// @param path 配置文件所在的路径
//
// @return error 错误信息
func LoadToml(path string) error {
	if _inited {
		return nil
	}

	// 1.加载: 配置文件
	err := loader.LoadToml(path, &_conf)

	// 2.校验: 环境参数
	if ok := stringz.ArrayContains(profiles, _conf.Env); !ok {
		return perrors.Errorf("config: the env config node candidate value is: %v", profiles)
	}

	// 3.配置文件 json 格式打印
	// 为什么需要打印？
	// 有时候可能想知道: app 解析的完成配置是什么?
	// 这个时候: 就可以将 属性: PrintConfig 设置为 true,
	// 这样 App 在启动的时候,就可以看到解析后配置信息
	if _conf.PrintConfig {
		printConfig()
	}

	// 4.标记: 已经初始化
	// 如果: 多次调用该方法: 第 N 次将不做任何副作用的操作(N > 1)
	_inited = true

	validateConfigItem()

	return err
}

// PrintConfig 当前 App 启动时,是否打印配置文件信息
func PrintConfig() bool {
	return _conf.PrintConfig
}

// Env 当前 App 启动的环境: env
func Env() string {
	return _conf.Env
}

// Host 当前 App 启动的 Host
func Host() string {
	return _conf.Host
}

func Timeout() uint32 {
	return _conf.TimeOut
}

// Log 日志配置信息
func Log() log.Config {
	return _conf.Log
}

// DatabaseMap 数据库配置信息
func DatabaseMap() map[string]DBConfig {
	return _conf.DatabaseMap
}

// DetermineDatabase 根据数据库: key 名称, 获取数据库配置信息
func DetermineDatabase(databaseKey string) DBConfig {
	return _conf.DatabaseMap[databaseKey]
}

// Rabbit 配置信息
func Rabbit() RabbitMQConfig {
	return _conf.RabbitMQ
}

// Redis 配置信息
func Redis() RedisConfig {
	return _conf.Redis
}

// printConfig json 打印配置文件信息
func printConfig() {
	pretty := jsonz.Pretty(_conf)
	fmt.Println(pretty)
}
