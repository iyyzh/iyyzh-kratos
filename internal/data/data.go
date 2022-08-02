package data

import (
	"iyyzh-kratos/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//4 基础设施层 业务数据访问 管理数据库初始化 实现了 biz 的 repo 接口 将领域对象重新拿出来

// ProviderSet is 依赖注入
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedis, NewUserRepo, NewOrderRepo)

// Data is 结构体 管理数据库连接
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewData is 注入实现数据库连接的 构造函数
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:  db,
		rdb: rdb,
	}, cleanup, nil
}

//NewDB is 连接mysql
func NewDB(c *conf.Data, logger log.Logger) *gorm.DB {
	//log := log.NewHelper(log.With(logger, "module", "realword-service/data/gorm"))

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	if err := db.AutoMigrate(); err != nil {
		log.Fatal(err)
	}

	return db
}

//NewRedis is 连接redis
func NewRedis(c *conf.Data, logger log.Logger) *redis.Client {
	//log := log.NewHelper(log.With(logger, "module", "realword-service/data/redis"))

	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})

	return rdb
}
