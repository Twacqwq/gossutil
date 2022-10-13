package config

import (
	"gossutil/pkg/helpers"
	"os"

	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
)

type ConfigFunc func() map[string]interface{}

var viper *viperlib.Viper
var ConfigFuncs map[string]ConfigFunc

func init() {
	viper = viperlib.New()
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("ossenv")
	viper.AutomaticEnv()
	ConfigFuncs = make(map[string]ConfigFunc)
}

// 初始化配置
func InitConfig(env string) {
	// 加载环境变量
	loadEnv(env)
	// 注册配置
	registerConfig()
}

// 添加配置
func Add(configName string, fn ConfigFunc) {
	ConfigFuncs[configName] = fn
}

// 读取环境变量
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue...)
	}
	return internalGet(envName)
}

// 获取String类型配置
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// 获取Int类型配置
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

func loadEnv(envSuffix string) {
	envPath := ".env"
	if len(envSuffix) > 0 {
		filePath := ".env." + envSuffix
		if _, err := os.Stat(filePath); err == nil {
			envPath = filePath
		}
	}
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.WatchConfig()
}

func registerConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	if !viper.IsSet(path) || helpers.IsEmpty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}
