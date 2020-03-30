package config

import (
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v2"
)

// MysqlConfig mysql配置
type MysqlConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

type InfuraConfig struct {
	Http string `yaml:"http"`
	Wss  string `yaml:"wss"`
}

type UsdtConfig struct {
	Token string `yaml:"token"`
}

type AddrConfig struct {
	Password string `yaml:"password"`
}

type Yaml struct {
	MysqlConfig  `yaml:"mysql"`
	InfuraConfig `yaml:"infura"`
	UsdtConfig   `yaml:"usdt"`
	AddrConfig   `yaml:"addr"`
}

var (
	conf Yaml
	once sync.Once
)

func GetConfig() Yaml {
	once.Do(func() {
		yamlFile, err := ioutil.ReadFile("./env.yaml")
		if err != nil {
			log.Fatalln(err)
		}
		if err = yaml.Unmarshal(yamlFile, &conf); err != nil {
			log.Fatalln(err)
		}
	})
	return conf
}
