package config

import (
	"github.com/athlon18/micro-core/config/structs"
	_const "github.com/athlon18/micro-core/const"
)

// Config 配置
type Config struct {
	Local      Mode `mapstructure:"local" json:"local" yaml:"local"`                // Local 本地
	Production Mode `mapstructure:"production" json:"production" yaml:"production"` // Production 生产
	K8s        Mode `mapstructure:"k8s" json:"k8s" yaml:"k8s"`                      // K8s 生产
}

type Mode struct {
	Mode               string                   `json:"mode"`
	App                structs.App              `mapstructure:"app" json:"app" yaml:"app"`
	Mysql              map[string]structs.Data  `json:"mysql" yaml:"mysql"`
	Redis              structs.Redis            `json:"redis" yaml:"redis"`
	Mongo              map[string]structs.Mongo `json:"mongo" yaml:"mongo"`
	Casbin             structs.Casbin           `json:"casbin" yaml:"casbin"`
	Consul             structs.Consul           `json:"consul" yaml:"consul"`
	Wechat             structs.Wechat           `json:"wechat" yaml:"wechat"`
	JWT                structs.JWT              `json:"jwt" yaml:"jwt"`
	CenterServerAddr   string                   `json:"CenterServerAddr" yaml:"CenterServerAddr"`
	BurgeonServiceAddr string                   `json:"BurgeonServiceAddr" yaml:"BurgeonServiceAddr"`
	RossServerAddr     string                   `json:"rossServerAddr" yaml:"rossServerAddr"`
	OaServerAddr       string                   `json:"oaServerAddr" yaml:"oaServerAddr"`
}

func (cfg Config) getLocal() Mode {
	return cfg.Local
}

func (cfg Config) getProduction() Mode {
	return cfg.Production
}

func (cfg Config) Get(mode string) Mode {
	switch mode {
	case _const.LOCAL:
		return cfg.getLocal().setMode(mode)
	case _const.PRODUCTION:
		return cfg.getProduction().setMode(mode)
	default:

	}
	return cfg.getLocal().setMode(_const.LOCAL)
}
func (mode Mode) setMode(str string) Mode {
	mode.Mode = str
	return mode
}

func (mode Mode) GetMysql() map[string]structs.Data {
	return mode.Mysql
}

//func (mode Mode) Mode() string {
//	return
//}
