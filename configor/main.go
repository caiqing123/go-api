package configor

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/mix-go/xcli/argv"
)

var Config = struct {
	Alipay struct {
		AppId      string
		PublicKey  string
		PrivateKey string
		NotifyUrl  string
		ReturnUrl  string
		IsProd     bool
	}
	Wechat struct {
		AppId      string
		MchId      string
		ApiKey     string
		SerialNo   string
		ApiV3Key   string
		PrivateKey string
		NotifyUrl  string
		ReturnUrl  string
		IsProd     bool
	}
}{}

func init() {
	// Conf support YAML, JSON, TOML, Shell Environment
	// auto reload configuration every second configor.New(&configor.Config{AutoReload: true})
	if err := configor.Load(&Config, fmt.Sprintf("%s/../conf/config.yml", argv.Program().Dir)); err != nil {
		panic(err)
	}
}
