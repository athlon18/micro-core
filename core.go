package core

import (
	"github.com/athlon18/micro-core/config"
	_const "github.com/athlon18/micro-core/const"
	"github.com/athlon18/micro-core/global"
	"github.com/athlon18/micro-core/utils"
	"github.com/athlon18/micro-core/viper"
	"github.com/gookit/color"
	"log"
	"os"
	"time"
)

type Core struct {
	Mode  config.Mode
	Utils utils.Util
}

func Initialize() *Core {

	// first
	if err := viper.Viper().Unmarshal(&global.Config); err != nil {
		log.Println("no config gotten")
	}

	// env APP_ENV select _const.LOCAL, _const.PRODUCTION
	mode := global.Config.Get(os.Getenv(_const.APP_ENV_NAME))

	color.Greenln("模式: "+mode.Mode+" 版本: "+mode.App.Ver+" 时间: "+time.Now().String(), "PodIP:", os.Getenv("GO_PODIP"))
	return &Core{
		Mode:  mode, // config
		Utils: utils.Util{},
	}
}
