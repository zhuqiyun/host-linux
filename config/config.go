package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var ParamsConfig *viper.Viper

func ParamsInit() {
	ParamsConfig = viper.New()
	ParamsConfig.AddConfigPath(".") //设置读取的文件路径
	ParamsConfig.SetConfigName("config")   //设置读取的文件名
	ParamsConfig.SetConfigType("yaml")     //设置文件的类型
	//尝试进行配置读取
	if err := ParamsConfig.ReadInConfig(); err != nil {
		fmt.Println("配置文件读取错误---：", err.Error())
	}
}
