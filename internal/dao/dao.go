package dao

import (
	"context"
	"time"

	"go-common/library/cache/memcache"
	"go-common/library/cache/redis"
	"go-common/library/conf/paladin.v2"
	"go-common/library/database/sql"
	"go-common/library/sync/pipeline/fanout"
	xtime "go-common/library/time"
	"helloword/internal/model"

	"github.com/google/wire"
)

var Provider = wire.NewSet(New, NewDB, NewRedis, NewMC)

type (
	//go:generate kratos tool btsgen
	// Dao dao interface
	Dao interface {
		Close()
		Ping(ctx context.Context) (err error)
		// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
		Article(c context.Context, id int64) (*model.Article, error)

		//新增接口
		AddUser(data *model.Users) (int64, error)
		UpdateUser(data *model.Users) (int64, error)
		DeleteUser(data *model.Users) (int64, error)
		SearchUser(data *model.Users) (map[string]string, error)
		SearchStructUser(data *model.Users) (*model.Users, error)
		SearchStruct() ([]*model.Users, error)

		RedisAdd(data *model.Users) (users string, err error)
		RedisDel(data *model.Users) (users string, err error)
		RedisGet(data *model.Users) (users string, err error)
		RedisUser(data *model.Users) (users *model.Users, err error)
		NewRedisGet(data *model.Users) (users *model.Users, err error)
	}
)

// dao dao.
type dao struct {
	db         *sql.DB
	redis      *redis.Redis
	mc         *memcache.Memcache
	cache      *fanout.Fanout
	demoExpire int32
}

// New new a dao and return.
func New(r *redis.Redis, mc *memcache.Memcache, db *sql.DB) (d Dao, cf func(), err error) {
	return newDao(r, mc, db)
}

// 根据参数初始化dao
func newDao(r *redis.Redis, mc *memcache.Memcache, db *sql.DB) (d *dao, cf func(), err error) {
	var cfg struct {
		DemoExpire xtime.Duration
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	d = &dao{
		db:         db,
		redis:      r,
		mc:         mc,
		cache:      fanout.New("cache"),
		demoExpire: int32(time.Duration(cfg.DemoExpire) / time.Second),
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *dao) Close() {
	d.cache.Close()
	// d.db.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
