package config

type Email struct {
	To          string `mapstructure:"to" json:"to" yaml:"to"`                               // 收件人:多个以英文逗号分隔
	Port        int    `mapstructure:"port" json:"port" yaml:"port"`                         // 发送端口
	From        string `mapstructure:"from" json:"from" yaml:"from"`                         // 收件人
	Host        string `mapstructure:"host" json:"host" yaml:"host"`                         // 服务器地址
	IsSSL       bool   `mapstructure:"is-ssl" json:"isSSL" yaml:"is-ssl"`                    // 是否SSL
	Secret      string `mapstructure:"secret" json:"secret" yaml:"secret"`                   // 密钥
	Nickname    string `mapstructure:"nickname" json:"nickname" yaml:"nickname"`             // 昵称
	ReceivePort int    `mapstructure:"receive-port" json:"receive-port" yaml:"receive-port"` // 接收端口
	ReceiveHost string `mapstructure:"receive-host" json:"receive-host" yaml:"receive-host"` // 服务器地址
}
