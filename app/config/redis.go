package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var cfgRedis *viper.Viper

var RedisConfig = new(Redis)

type Redis struct {
	Addr     string
	Port     int
	Password string
	DB       int
}

//载入配置文件
func SetupRedisConfig(path string) {
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}
	cfgRedis = viper.Sub("config.redis")
	if cfgRedis == nil {
		panic("No found config.redis in the configuration")
	}
	RedisConfig = initRedis(cfgRedis)
}

func initRedis(cfg *viper.Viper) *Redis {
	config := &Redis{
		Addr:     cfg.GetString("addr"),
		Port:     cfg.GetInt("port"),
		Password: cfg.GetString("password"),
		DB:       cfg.GetInt("db"),
	}
	return config
}

// SetupRedisConfigNew 载入配置文件
func SetupRedisConfigNew(addr string, port int, password string, db int) {
	RedisConfig = &Redis{
		Addr:     addr,
		Port:     port,
		Password: password,
		DB:       db,
	}
}
