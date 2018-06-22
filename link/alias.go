package link

import (
	"Go-DMT/Go-DTM/link/dblink"
	"Go-DMT/Go-DTM/link/rabbitmq-dirver"
	"database/sql"
	"github.com/Sirupsen/logrus"
	"github.com/astaxie/beego/logs"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

type DriverType int

const (
	_ DriverType = iota
	DRMySQL
	RabbitMQ
	DBRedis
)

//写一个driver的模型，用户获取不同驱动
type driver string

func (d driver) Type() DriverType {
	a, _ := dataBaseCache.get(string(d))
	return a.Driver
}

func (d driver) Name() string {
	return string(d)
}

// check driver iis implemented Driver interface or not.
//var _ Driver = new(driver)

var (
	dataBaseCache = &_dbCache{cache: make(map[string]*alias)}
	drivers       = map[string]DriverType{
		"mysql":    DRMySQL,
		"rabbitmq": RabbitMQ,
	}
	dbBasers = map[DriverType]dbBase{
		DRMySQL: dblink.NewdbBaseMysql(),
		//TODO rabbitMQ的操作类

	}
)

//主要用于操作alias的缓存
type _dbCache struct {
	mu    sync.Mutex
	cache map[string]*alias
}

func (ac *_dbCache) get(name string) (al *alias, ok error) {
	ac.mu.Lock()
	defer ac.mu.Unlock()
	al, ok = ac.cache[name]
	return
}

//TODO add()
func (ac *_dbCache) add(name string, al *alias) (added bool) {
	ac.mu.Lock()
	defer ac.mu.Unlock()
	if _, ok := ac.cache[name]; !ok {
		ac.cache[name] = al
		added = true
	}
	return
}

//TODO get default

//TODO 可能会根据rabbitMQ去增加一些特定的选项
type alias struct {
	Name         string
	Driver       DriverType
	DriverName   string
	DataSource   string
	MaxIdleConns int
	MaxOpenConns int
	DB           interface{}
	DbBaser      dbBase
	TZ           *time.Location
	Engine       string
}

//TODO设置数据库的时间，如果没有设置，默认用系统时间

func defaultDZ(a *alias) {
	a.TZ = time.Local

	switch a.Driver {
	case DRMySQL:
		//mysql数据库时间设置
		//mysql数据库引擎设置
		//mysql数据库的其他管理参数设置，比如数据库编码类型等

	//TODO 如果其他的数据库也需要数据设置时间，可以在这里增加，并且在常量中添加
	default:

	}
}

//TODO 每个存储介质的连接都需要经过三个函数，addAliasWithDB,AddAliasWithDB,RegisterDataBase。

func addAliasWithDB(aliasName string, driverName string, db interface{}) (al *alias) {
	al = new(alias)
	al.Name = aliasName
	al.DB = db

	if dr, ok := drivers[driverName]; ok {
		al.Driver = dr
		al.DbBaser = dbBasers[dr]
	}
	if !dataBaseCache.add(aliasName, al) {
		logs.Debug("the alias cache has been")
	}
	return
}

func AddAliasWithDB(aliasName string, driverName string, db interface{}) (al *alias) {
	al = addAliasWithDB(aliasName, driverName, db)
	return
}

//TODO 具体实现连接
func RegisterDataBase(aliasName string, driverName string, dbRoute string, params ...int) error {
	//TODO 由于database的sql连接池没办法满足rabbitmq的连接使用，只能在这里做个分支，可以考虑一下有什么更好设计模式方案
	var (
		err error
		db  interface{}
	)

	if driverName == "rabbitmq" {
		db, err = rabbitmq_dirver.Open(driverName, dbRoute)
	} else {
		db, err = sql.Open(driverName, dbRoute)
	}
	//db, err := sql.Open(driverName, dbRoute)

	if err != nil {
		logs.Debug("connect db fail")
	}

	al := AddAliasWithDB(aliasName, driverName, db)

	al.DataSource = dbRoute

	for k, v := range params {
		switch k {
		case 0:
			//TODO 设置最大连接缓存队列，因为sql的open用了封装了连接池
			SetMaxIdleConns(aliasName, v)
		case 1:
			SetMaxOpenConns(aliasName, v)

		default:

		}
	}
	return nil
}

func RegisterDriver(driverName string, typ DriverType) error {
	if t, ok := drivers[driverName]; !ok {
		drivers[driverName] = typ
	} else {
		if t != typ {
			logrus.Debug("driver has register but it is not the driverType")
		}
	}
	return nil

}

//以下是辅助方法
func SetDataBaseTZ() {

}

func SetMaxOpenConns(alias string, maxConns int) error {
	al, err := dataBaseCache.get(alias)
	if err != nil {
		logs.Debug("get alias cache failed")
		return err
	}
	al.MaxOpenConns = maxConns
	al.DB.SetMaxOpenConns(maxConns)
	return nil
}

func SetMaxIdleConns(aliasName string, maxIdelNum int) error {
	al, err := dataBaseCache.get(aliasName)
	if err != nil {
		return err
	}
	al.MaxIdleConns = maxIdelNum
	al.DB.SetMaxIdleConns(maxIdelNum)
	return nil
}

func GetDB(aliasName string) *sql.DB {
	al, err := dataBaseCache.get(aliasName)
	if err != nil {
		logs.Debug("get alias cache failed")
		return nil
	}
	return al.DB
}
