package structs

type Mysql struct {
	Center Data `json:"center" yaml:"center"`
}

type Data struct {
	Dsn   string `json:"dsn" yaml:"dsn"`
	Debug bool   `json:"debug" yaml:"debug"`
}
