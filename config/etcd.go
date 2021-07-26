package config

import (
	"time"

	"github.com/yangliang4488/log_agent_system/logger"
	"go.etcd.io/etcd/clientv3"
)

type Etcd struct {
	Address    string
	LogKey     string
	SysinfoKey string
}

var (
	client *clientv3.Client
)

func ConnEtcd(address []string) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   address,
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	return
}
