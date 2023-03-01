package testdata

import (
	"github.com/li-zeyuan/CSEIdiom/backend/config"
	"github.com/li-zeyuan/CSEIdiom/backend/dao"
	"github.com/li-zeyuan/common/mylogger"
	"github.com/li-zeyuan/common/mysqlstore"
)

func InitServer() error {
	cfg := &config.Config{
		Mysql:mysqlstore.Config{
			DSN:     "root:root@tcp(localhost:3306)/government_exam?charset=utf8mb4&parseTime=True&loc=UTC",
			MaxConn: 1,
			MaxOpen: 1,
			Timeout: 10,
		},
		Logging:mylogger.LoggerCfg{

		},
	}
	config.AppCfg = cfg
	mylogger.Init(nil)
	dao.New(cfg)

	return mysqlstore.New(&cfg.Mysql)
}
