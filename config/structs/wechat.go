package structs

type Wechat struct {
	Url  string                `json:"url" yaml:"url"`
	Data map[string]WechatData `json:"data" yaml:"data"`
}

type WechatData struct {
	Webhook string `json:"webhook" yaml:"webhook"`
}
