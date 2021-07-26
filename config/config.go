package config

import (
	"fmt"

	"github.com/yangliang4488/log_agent_system/logger"

	"github.com/spf13/viper"
)

type ConfigClass struct {
	Kfaka
	Etcd
}

var Config *ConfigClass

func LoadConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(".env 配置文件未找到")
			return
		}
	} else {
		kafkaConf := Kfaka{
			Address:  viper.GetString("KAFKA_ADDRESS"),
			ChanSize: viper.GetInt("KAFKA_CHAN_SIZE"),
		}
		etcdConf := Etcd{
			Address:    viper.GetString("ETCD_ADDRESS"),
			LogKey:     viper.GetString("ETCD_COLLECT_LOG_KEY"),
			SysinfoKey: viper.GetString("ETCD_COLLECT_SYSINFO_KEY"),
		}
		Config = &ConfigClass{
			Kfaka: kafkaConf,
			Etcd:  etcdConf,
		}

		logger.Logger.Info("load configure success")
	}
}
