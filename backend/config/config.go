package config

import (
	"io/ioutil"
	"log"

	"github.com/li-zeyuan/common/mylogger"
	"github.com/li-zeyuan/common/mysqlstore"
	"gopkg.in/yaml.v3"
)

var AppCfg *Config

type Config struct {
	ListenAddress string             `yaml:"listen_address"`
	Timeout       int                `yaml:"timeout"`
	JwtSecret     string             `yaml:"jwt_secret"`
	WxAppId       string             `yaml:"wx_app_id"`
	WxSecret      string             `yaml:"wx_secret"`
	Mysql         mysqlstore.Config  `yaml:"mysql"`
	Logging       mylogger.LoggerCfg `yaml:"logging"`
}

func LoadConfigFile(cfgPath string) error {
	buf, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		log.Printf("Read config file(%s) failed: %s.\n", cfgPath, err)
		return err
	}

	var conf Config
	if err = yaml.Unmarshal(buf, &conf); err != nil {
		log.Printf("Unmarshal config file(%s) failed: %s.\n", cfgPath, err)
		return err
	}

	log.Println("tracing_query config: ", conf)
	cfg := withDefaultCfg(conf)
	err = verifyCfg(cfg)
	if err != nil {
		return err
	}

	AppCfg = cfg

	return nil
}

func withDefaultCfg(cfg Config) *Config {
	cfg.Timeout = 3
	log.Println("tracing_query with default config: ", cfg)
	return &cfg
}

func verifyCfg(cfg *Config) error {

	return nil
}
