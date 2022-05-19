package structs

type App struct {
	Ver      string `json:"ver" yaml:"ver"`
	Prefix   string `json:"prefix" yaml:"prefix"`
	Mode     string `json:"mode" yaml:"mode"`
	HttpPort string `mapstructure:"http-port" json:"http_port" yaml:"http-port"`
}
