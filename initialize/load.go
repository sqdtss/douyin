package initialize

import (
	"douyin/global"
	"fmt"
	"github.com/spf13/viper"
)

// LoadConfig 加载配置文件
func LoadConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("viper readinconfig error: %s\n", err))
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("viper unmarshal err: %s\n", err))
	}
}
