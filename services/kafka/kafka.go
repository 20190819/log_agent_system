package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/yangliang4488/log_agent_system/config"
	"github.com/yangliang4488/log_agent_system/logger"
)

var msgChan chan *Message

type Message struct {
	Data  string
	Topic string
}

func SendMsgToChan(msg *Message) (err error) {
	select {
	case msgChan <- msg:
	default:
		err = fmt.Errorf("kafaka msgChan is empty")
	}
	return
}

func SendToKafka() {
	defer config.Producer.Close()
	for msg := range msgChan {
		kfmsg := &sarama.ProducerMessage{}
		kfmsg.Topic = msg.Topic
		kfmsg.Value = sarama.StringEncoder(msg.Data)
		pid, offset, err := config.Producer.SendMessage(kfmsg)
		if err != nil {
			logger.Logger.Errorf("send to kafka msg err: ", err)
			continue
		}
		logger.Logger.Infof("send msg success, pid:%v offset:%v\n", pid, offset)
	}
}
