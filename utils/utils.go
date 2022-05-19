package utils

var (
	ConfigEnv  = "APP_CONFIG"
	ConfigFile = "config.yaml"
)

type Util struct {
}

func SetConfigEnv(str string) {
	ConfigEnv = str
}

func SetConfigFile(str string) {
	ConfigFile = str
}
