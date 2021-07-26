package bootstrap

import (
	"github.com/yangliang4488/log_agent_system/config"
	"github.com/yangliang4488/log_agent_system/logger"
	"strings"
)

func App() {
	config.LoadConfig()
	err := config.ConnKafka(strings.Split(config.Config.Kfaka.Address, ","))
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	err = config.ConnEtcd(strings.Split(config.Config.Etcd.Address, ","))
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
}
