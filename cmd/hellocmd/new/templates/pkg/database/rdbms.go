package database

import (
	"fmt"

	"codeup.aliyun.com/uphicoo/gokit/pkg/database/rdbms/driverproxy"
	"codeup.aliyun.com/uphicoo/gokit/pkg/database/rdbms/driverwrapper"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent"
	_ "github.com/go-sql-driver/mysql" // mysql 驱动
	_ "github.com/jackc/pgx/v4/stdlib" // postgresql 驱动
	perrors "github.com/pkg/errors"

	"uphicoo.com/uphicoo/project-template/internal/config"
)

var (
	_            dialect.Driver = (*driverwrapper.DriverWrapper)(nil)
	_clientProxy *ent.Client    // notices | merchants 数据库客户端代理
)

// RDBMSClientInit 初始化-据库客户端
func RDBMSClientInit(confMap map[string]config.DBConfig) error {
	confN, ok := confMap[NoticeDatabase]
	if !ok {
		return perrors.Errorf("database.rdbms: 数据库:[%s]配置不存在,请核实", NoticeDatabase)
	}

	err := initDBClientProxy(confN)
	if err != nil {
		return perrors.Errorf("database.rdbms: 数据库: 配置数据库驱动代理失败:%v", err)
	}

	return nil
}

// RDBMSClientClose 关闭-据库客户端
func RDBMSClientClose() error {
	err := _clientProxy.Close()
	if err != nil {
		return perrors.New("database.clientProxy:close.error" + err.Error())
	}

	return nil
}

// ---------------------------------------------------------------- clients (temp design)

// ProxyClient 返回 notices | merchants 数据库客户端 代理
func ProxyClient() *ent.Client {
	return _clientProxy
}

// ---------------------------------------------------------------- private

func initDBClientProxy(confN config.DBConfig) error {
	// 初始化 notices 数据库驱动
	driverN, err := initNoticeDriver(confN)
	if err != nil {
		return err
	}

	driverProxy := driverproxy.NewDriverProxy()
	driverProxy.RegisterDialect(driverN.Dialect())

	if driverN != nil {
		err = driverProxy.RegisterDriver(NoticeDatabase, driverN)
		if err != nil {
			return err
		}
	}

	_clientProxy = ent.NewClient(ent.Driver(driverProxy))

	return nil
}

func initNoticeDriver(confN config.DBConfig) (dialect.Driver, error) {
	// 不实例化
	if !confN.Enabled {
		return nil, nil
	}

	writeDb, err := determineDbDriver(confN, true)
	if err != nil {
		return nil, err
	}
	readDb, err := determineDbDriver(confN, false)
	if err != nil {
		return nil, err
	}

	driverN := driverwrapper.NewDriverWrapper(writeDb, readDb)
	driverN.RegisterDialect(confN.Driver)

	return driverN, nil
}

func initMerchantDriver(confM config.DBConfig) (dialect.Driver, error) {
	// 不实例化
	if !confM.Enabled {
		return nil, nil
	}

	writeDb, err := determineDbDriver(confM, true)
	if err != nil {
		return nil, err
	}
	readDb, err := determineDbDriver(confM, false)
	if err != nil {
		return nil, err
	}

	driverM := driverwrapper.NewDriverWrapper(writeDb, readDb)
	driverM.RegisterDialect(confM.Driver)

	return driverM, nil
}

func determineDbDriver(conf config.DBConfig, writeHost bool) (*sql.Driver, error) {
	var host string
	var port int
	var db *sql.Driver
	var err error

	if writeHost {
		host = conf.WriteHost
		port = conf.WritePort
	} else {
		host = conf.ReadHost
		port = conf.ReadPort
	}

	switch conf.Driver {
	case dialect.Postgres:
		db, err = populatePostgres(conf, db, err, host, port)
	case dialect.MySQL:
		db, err = populateMySQL(conf, db, err, host, port)
	default:
		return nil, perrors.Errorf("database: 不受支持的数据库驱动:%s", conf.Driver)
	}
	if err != nil {
		return nil, err
	}

	initConnectionParams(conf, db)

	return db, nil
}

func initConnectionParams(conf config.DBConfig, db *sql.Driver) {
	sqlDB := db.DB()
	sqlDB.SetConnMaxIdleTime(conf.ConnMaxIdleTime.Duration)
	sqlDB.SetConnMaxLifetime(conf.ConnMaxLifeTime.Duration)
	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
}

func populateMySQL(conf config.DBConfig, db *sql.Driver, err error, host string, port int) (*sql.Driver, error) {
	db, err = sql.Open(
		dialect.MySQL,
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=True&charset=utf8&loc=%s",
			conf.Username,
			conf.Password,
			host,
			port,
			conf.Dbname,
			"Asia%2FShanghai",
		),
	)

	return db, err
}

func populatePostgres(conf config.DBConfig, db *sql.Driver, err error, host string, port int) (*sql.Driver, error) {
	db, err = sql.Open(
		dialect.Postgres,
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s?TimeZone=%s",
			conf.Username,
			conf.Password,
			host,
			port,
			conf.Dbname,
			"Asia%2FShanghai",
		),
	)

	return db, err
}