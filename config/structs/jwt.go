package structs

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signing_key" yaml:"signing-key"`
	ExpiresTime int64  `mapstructure:"expires-time" json:"expires_time" yaml:"expires-time"`
	BufferTime  int64  `mapstructure:"buffer-time" json:"buffer_time" yaml:"buffer-time"`
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
}
