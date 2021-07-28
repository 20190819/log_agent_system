package tailfile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/yangliang4488/log_agent_system/logger"
	kafkaService "github.com/yangliang4488/log_agent_system/services/kafka"
	"github.com/yangliang4488/log_agent_system/utils"
)

var localIP string

type LogData struct {
	IP   string `json:"ip"`
	Data string `json:"data"`
}

func init() {
	localIP, _ = utils.GetLocalIP()
}

type tailClass struct {
	path    string
	module  string
	topic   string
	instace *tail.Tail
	ctx     context.Context
	cancel  context.CancelFunc
}

func (t *tailClass) Init() (err error) {
	t.instace, err = tail.TailFile(t.path, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("init tail failed: ", err.Error())
	}
	return
}

// 每个 tailClass 都要单独读取日志信息发送到 kafka 中
func (t *tailClass) run() {
	for {
		select {
		case <-t.ctx.Done():
			logger.Logger.Warningf("the task for path: %s is stop .", t.path)
			t.instace.Cleanup()
			return
		case line, ok := <-t.instace.Lines:
			if !ok {
				logger.Logger.Error("read line failed")
				continue
			}

			logData := &LogData{
				IP:   localIP,
				Data: line.Text,
			}
			jsonStr, err := json.Marshal(logData)
			if err != nil {
				logger.Logger.Error("json 序列化失败", err)
				continue
			}

			msg := &kafkaService.Message{
				Data:  string(jsonStr),
				Topic: t.topic, // 先写死
			}

			err = kafkaService.SendMsgToChan(msg)
			if err != nil {
				logger.Logger.Error(err)
			}
		}
	}
}

func CreateTailObject(path string, module string, topic string) (t *tailClass, err error) {
	t = &tailClass{
		path:   path,
		module: module,
		topic:  topic,
	}
	t.ctx, t.cancel = context.WithCancel(context.Background())
	// 初始化
	err = t.Init()
	return
}
