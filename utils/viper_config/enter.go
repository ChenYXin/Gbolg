package viper_config

import (
	"fmt"
	"path"
	"runtime"

	"github.com/spf13/viper" // 相对于$GOPATH/pkg/mod的路径
)

var (
	ProjectRootPath = path.Dir(getoncurrentPath()+"/../../") + "/"
	VC              *viper.Viper
)

func getoncurrentPath() string {
	_, filename, _, _ := runtime.Caller(0) //0表示当前本行代码在什么位置
	return path.Dir(filename)              //返回文件所在的目录
}

// InitViperConfig Viper可以解析JSON、TOML、YAML、HCL、INI、ENV等格式的配置文件。甚至可以监听配置文件的变化(WatchConfig)，不需要重启程序就可以读到最新的值。
func InitViperConfig() {
	VC = viper.New()
	configPath := ProjectRootPath + "conf/"
	VC.AddConfigPath(configPath) // 文件所在目录
	VC.SetConfigName("app")      // 文件名
	VC.SetConfigType("yaml")     // 文件类型
	configFile := configPath + "app.yaml"

	if err := VC.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("找不到配置文件:%s", configFile)) //系统初始化阶段发生任何错误，直接结束进程
		} else {
			panic(fmt.Errorf("解析配置文件%s出错:%s", configFile, err))
		}
	}
}
