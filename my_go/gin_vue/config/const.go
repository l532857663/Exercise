package config

// 订单循环监控配置
type OrderPaymentConf struct {
	Enable        bool  `mapstructure:"enable" json:"enable" yaml:"enable"`                         // 是否开启服务
	Time          int64 `mapstructure:"time" json:"time" yaml:"time"`                               // 订单待支付时间（秒）
	CheckNum      int   `mapstructure:"check-num" json:"check-num" yaml:"check-num"`                // 检测订单状态的次数
	CheckInterval int64 `mapstructure:"check-interval" json:"check-interval" yaml:"check-interval"` // 检测间隔时间(秒)
	TempInterval  int64 `mapstructure:"temp-interval" json:"temp-interval" yaml:"temp-interval"`    // 循环队列的步长(秒)
	Temp          int64 `mapstructure:"temp" json:"temp" yaml:"temp"`                               // 首次添加订单的延后步数
}

// 常量参数配置
type ConstData struct { // mapstructure:"" json:"" yaml:""
	EmailTitle          string `mapstructure:"email-title" json:"email-title" yaml:"email-title"`                                  // 发送的邮件标题
	RegisterUrl         string `mapstructure:"register-url" json:"register-url" yaml:"register-url"`                               // 注册地址的邮件链接
	ForgetUrl           string `mapstructure:"forget-url" json:"forget-url" yaml:"forget-url"`                                     // 忘记密码的邮件链接
	EmailResendTime     int    `mapstructure:"email-resend-time" json:"email-resend-time" yaml:"email-resend-time"`                // 邮件重发的间隔时间(秒)
	EmailResendNum      int64  `mapstructure:"email-resend-num" json:"email-resend-num" yaml:"email-resend-num"`                   // 邮件重发次数
	EmailResendNumTime  int    `mapstructure:"email-resend-num-time" json:"email-resend-num-time" yaml:"email-resend-num-time"`    // 邮件重发次数的有效期(秒)
	EmailVerifySaveTime int    `mapstructure:"email-verify-save-time" json:"email-verify-save-time" yaml:"email-verify-save-time"` // 验证邮件有效期(秒)
	OrderMintTime       string `mapstructure:"order-mint-time" json:"order-mint-time" yaml:"order-mint-time"`                      // token的mint时间 eg: 2021-08-31 10:00:00
}
