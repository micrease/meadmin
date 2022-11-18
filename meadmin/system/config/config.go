package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"meadmin/library/files"
	rpcx "meadmin/library/rpc/config"
	"path"
)

const (
	ConfigDir        = "./resources/"
	ConfigFileFormat = "config-%s.yaml"
)

const (
	EnvProd = "prod"
	EnvTest = "test"
	EnvDev  = "dev"
)

const (
	ModeDebug   = "debug"
	ModeRelease = "release"
)

type Database struct {
	Dsn string `yaml:"dsn"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

// 服务端配置
type ServiceConfig struct {
	ServiceName string      `yaml:"service_name"`
	Port        string      `yaml:"port"`
	Mode        string      `yaml:"mode"`
	Env         string      `yaml:"env"`
	Version     string      `yaml:"version"`
	KafkaBroker string      `yaml:"kafka_broker"`
	DocEnable   bool        `yaml:"doc_enable"`
	Database    Database    `yaml:"database"`
	JwtSecret   string      `yaml:"jwt_secret"`
	Redis       Redis       `yaml:"redis"`
	Rpcx        rpcx.Config `yaml:"rpcx"`
	ProjectPath string
}

func (config *ServiceConfig) IsProdEnv() bool {
	if config.Env != EnvProd {
		return false
	}
	return true
}

func (config *ServiceConfig) IsDebug() bool {
	if config.IsProdEnv() {
		return false
	}
	if config.Mode == ModeDebug {
		return true
	}
	return false
}

func (config *ServiceConfig) IsDocEnable() bool {
	if config.IsProdEnv() {
		return false
	}
	return config.DocEnable
}

var config = new(ServiceConfig)

func GetConfig() *ServiceConfig {
	return config
}

func GetConfigPath(env string) string {
	return path.Join(ConfigDir, fmt.Sprintf(ConfigFileFormat, env))
}

func InitConfig(path string) *ServiceConfig {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("InitConfig: ", path, ".  config err: ", err.Error())
		return nil
	}

	if err = yaml.Unmarshal(content, config); err != nil {
		log.Fatal("InitConfig: ", path, ".  config err: ", err.Error())
		return nil
	}
	config.ProjectPath = files.ProjectDir()
	return config
}
