package config

import (
	"github.com/li-zeyuan/common/mylogger"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var AppCfg *Config

type Config struct {
	ListenAddress string             `yaml:"listen_address"`
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

	return nil
}

func withDefaultCfg(cfg Config) *Config {
	log.Println("tracing_query with default config: ", cfg)
	return &cfg
}

func verifyCfg(cfg *Config) error {

	return nil
}
