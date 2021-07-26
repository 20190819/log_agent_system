package config

import (
	"github.com/Shopify/sarama"
	"github.com/yangliang4488/log_agent_system/logger"
	"time"
)

var (
	Producer sarama.SyncProducer
)

type Kfaka struct {
	Address  string
	ChanSize int
}

func ConnKafka(address []string) (err error) {
	c := sarama.NewConfig()

	c.Producer.MaxMessageBytes = 1000000
	c.Producer.RequiredAcks = sarama.WaitForLocal
	c.Producer.Timeout = 3 * time.Second
	c.Producer.Partitioner = sarama.NewHashPartitioner
	c.Producer.Retry.Max = 3
	c.Producer.Retry.Backoff = 100 * time.Millisecond
	c.Producer.Return.Successes = true

	Producer, err = sarama.NewSyncProducer(address, c)
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}

	return
}
