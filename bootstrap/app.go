package bootstrap

import (
	"github.com/yangliang4488/log_agent_system/config"
	"github.com/yangliang4488/log_agent_system/logger"
)

func App() {
	config.LoadConfig()
	err := config.ConnKafka([]string{config.Config.Kfaka.Address}, config.Config.Kfaka.ChanSize)
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	err = config.ConnEtcd([]string{config.Config.Etcd.Address})
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
}
