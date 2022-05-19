package structs

type Redis struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"`
}
