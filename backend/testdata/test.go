package testdata

import (
	"github.com/li-zeyuan/CSEIdiom/backend/config"
	"github.com/li-zeyuan/common/mylogger"
	"github.com/li-zeyuan/common/mysqlstore"
)

func InitServer() error {
	err := config.LoadConfigFile("/Users/zeyuan.li/Desktop/workspace/code/src/github.com/li-zeyuan/sun/governmentexam/backend_config_dev.yaml")
	if err != nil {
		return err
	}

	mylogger.Init(&config.AppCfg.Logging)

	return mysqlstore.New(&config.AppCfg.Mysql)
}
