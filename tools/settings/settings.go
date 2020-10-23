package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	*ServerConfig `mapstructure:"server"`
	*LogConfig    `mapstructure:"log"`
	*MysqlConfig  `mapstructure:"mysql"`
	*RedisConfig  `mapstructure:"redis"`
}
type ServerConfig struct {
	*HttpConfig `mapstructure:"http"`
}

type HttpConfig struct {
	Port int `mapstructure:"port"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_Size"`
}

func Init() (err error) {
	viper.SetConfigFile("./config/server.yaml") // 指定配置文件
	//viper.AddConfigPath("./conf/")     // 指定查找配置文件的路径
	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed,err:%v\n", err) // 读取配置信息失败
		return err
	}
	//把读取到的信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper Unmarshal faild,err:%v\n", err)
	}
	// 监控配置文件变化(配置热加载)
	viper.WatchConfig()
	//config文件修改后的回调函数
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper Unmarshal faild,err:%v\n", err)
		}
	})
	return err
}
