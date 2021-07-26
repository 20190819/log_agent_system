package etcd

import (
	"context"
	"encoding/json"
	"github.com/yangliang4488/log_agent_system/config"
	"github.com/yangliang4488/log_agent_system/logger"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var confChan = make(chan []*CollectEntry)

type CollectEntry struct {
	Path   string `json:"path"`
	Module string `json:"module"`
	Topic  string `json:"topic"`
}

type CollectSystemInfoEntry struct {
	Interval int64  `json:"interval"`
	Topic    string `json:"topic"`
}

func GetCollectConf(key string) (confCollect []*CollectEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := config.ClientEtcd.Get(ctx, key)
	if err != nil {
		logger.Logger.Error("从 etcd 获取收集日志配置失败", err.Error())
		return
	}

	if len(response.Kvs) == 0 {
		logger.Logger.Error("从 etcd 获取值为空")
		return
	}

	keyValue := response.Kvs[0]

	err = json.Unmarshal(keyValue.Value, &confCollect)
	if err != nil {
		logger.Logger.Error("json unmarshal failed: ", err.Error())
		return
	}
	logger.Logger.Debugf("load collect conf from etcd success %-v", confCollect)
	return
}

func GetSystemInfoConf(key string) (confSystemInfo []*CollectSystemInfoEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := config.ClientEtcd.Get(ctx, key)
	if err != nil {
		logger.Logger.Error("从 etcd 获取系统信息配置失败", err.Error())
		return
	}

	if len(response.Kvs) == 0 {
		logger.Logger.Error("从 etcd 获取值为空")
		return
	}

	keyValue := response.Kvs[0]

	err = json.Unmarshal(keyValue.Value, &confSystemInfo)
	if err != nil {
		logger.Logger.Error("json unmarshal failed: ", err.Error())
		return
	}
	logger.Logger.Debugf("load collect conf from etcd success %-v", confSystemInfo)
	return
}

func WatchConfChan() <-chan []*CollectEntry {
	return confChan
}

func WatchConf(key string) {
	for {
		rch := config.ClientEtcd.Watch(context.Background(), key) // <-chan WatchResponse
		for wresp := range rch {
			if err := wresp.Err(); err != nil {
				logger.Logger.Warnf("watch key:%s err:%v", key, err)
				continue
			}
			for _, ev := range wresp.Events {
				// 获取了最新的日志配置项怎么传给tailTaskMgr呢？
				var newConf []*CollectEntry
				// 需要判断如果是删除操作
				if ev.Type == clientv3.EventTypeDelete {
					confChan <- newConf
					continue
				}
				err := json.Unmarshal(ev.Kv.Value, &newConf)
				if err != nil {
					logger.Logger.Warnf("unmarshal the conf from etcd failed, err:%v", err)
					continue
				}
				confChan <- newConf
			}
		}
	}
}
