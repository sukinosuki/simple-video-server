package config

import "github.com/BurntSushi/toml"

type envConfig struct {
	Name  string
	Port  int
	Debug bool
}

type mysqlConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type redisConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database int
}

type jwtConfig struct {
	SecretKey     string
	EffectiveTime int
}

type logConfig struct {
	MaxAge     int
	MaxSize    int
	MaxBackups int
}

type ossConfig struct {
	Protocol        string
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}

type appConfig struct {
	Env   envConfig
	Mysql mysqlConfig
	Jwt   jwtConfig
	Redis redisConfig
	Oss   ossConfig
	Log   logConfig
}

var (
	Env   envConfig
	Mysql mysqlConfig
	Jwt   jwtConfig
	Redis redisConfig
	Oss   ossConfig
	Log   logConfig
)

func init() {
	var _appConfig appConfig

	_, err := toml.DecodeFile("config/config.toml", &_appConfig)

	if err != nil {
		panic(err)
	}

	Env = _appConfig.Env

	Mysql = _appConfig.Mysql

	Jwt = _appConfig.Jwt

	Redis = _appConfig.Redis

	Oss = _appConfig.Oss

	Log = _appConfig.Log
}
