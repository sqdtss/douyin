package config

// Config 组合全部配置模型
type Config struct {
	Server Server `mapstructure:"server"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Redis  Redis  `mapstructure:"redis"`
	Upload Upload `mapstructure:"upload"`
	Jwt    Jwt    `mapstructure:"jwt"`
}

// Server 服务启动端口号配置
type Server struct {
	Protocol string `mapstructure:"protocol"`
	Port     string `mapstructure:"port"`
}

// Mysql MySQL数据源配置
type Mysql struct {
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Url             string `mapstructure:"url"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

// Redis Redis配置
type Redis struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// Upload 文件上传相关路径配置
type Upload struct {
	SavePath  string `mapstructure:"savePath"`
	AccessUrl string `mapstructure:"accessUrl"`
}

// Jwt 用户认证配置
type Jwt struct {
	SigningKey string `mapstructure:"signingKey"`
}
