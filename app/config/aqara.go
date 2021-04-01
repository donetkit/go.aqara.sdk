package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var aqaraApp *viper.Viper

var AqaraConfig = new(Aqara)

type Aqara struct {
	OauthClientId     string `json:"oauth_client_id"`     // 第三方应用ID，AppID
	OauthClientSecret string `json:"oauth_client_secret"` // 第三方应用秘钥，AppKey
	DevClientId       string `json:"dev_client_id"`       // 第三方应用ID，AppID
	DevClientSecret   string `json:"dev_client_secret"`   // 第三方应用秘钥，AppKey
	AuthorizeUri      string `json:"authorize_uri"`       // 回调地址
}

func initAqara(cfg *viper.Viper) *Aqara {
	return &Aqara{
		OauthClientId:     cfg.GetString("oauth_client_id"),
		OauthClientSecret: cfg.GetString("oauth_client_secret"),
		DevClientId:       cfg.GetString("dev_client_id"),
		DevClientSecret:   cfg.GetString("dev_client_secret"),
		AuthorizeUri:      cfg.GetString("authorize_uri"),
	}
}

//载入配置文件
func SetupAqaraConfig(path string) {
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}
	aqaraApp = viper.Sub("config.aqara")
	if aqaraApp == nil {
		panic("No found config.aqara in the configuration")
	}
	AqaraConfig = initAqara(aqaraApp)
}
