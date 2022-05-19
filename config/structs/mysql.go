package structs

type Mysql struct {
	Center Center `json:"center" yaml:"center"`
}

type Center struct {
	Dsn   string `json:"dsn" yaml:"dsn"`
	Debug bool   `json:"debug" yaml:"debug"`
}
