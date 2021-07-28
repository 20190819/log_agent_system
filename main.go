package main

import (
	"sync"

	"github.com/yangliang4488/log_agent_system/bootstrap"
	"github.com/yangliang4488/log_agent_system/config"
	"github.com/yangliang4488/log_agent_system/logger"
	etcdService "github.com/yangliang4488/log_agent_system/services/etcd"
	"github.com/yangliang4488/log_agent_system/services/kafka"
	"github.com/yangliang4488/log_agent_system/services/tailfile"
)

var (
	wg = sync.WaitGroup{}
)

func main() {
	// 载入基础配置
	bootstrap.App()
	// goroutine 发送消息到 kafka
	go kafka.SendToKafka()
	// 从 etcd 获取【收集日志】配置信息
	collectLogKey := config.Config.Etcd.LogKey
	collectLogConf, _ := etcdService.GetCollectConf(collectLogKey)
	// 监听配置通道
	newConfChan := etcdService.WatchConfChan()
	// 初始化 tailManger
	err := tailfile.InitManager(collectLogConf, newConfChan)
	if err != nil {
		panic(err.Error())
	}
	logger.Logger.Info("init tail success .")
	// 执行
	run(collectLogKey)
	logger.Logger.Info("log agent over .")
}

// 业务逻辑处理
func run(logConfKey string) {
	// 实时监控etcd中日志收集配置项的变化，对tailObj进行管理
	wg.Add(2)
	go etcdService.WatchConf(logConfKey)
	wg.Wait()
}
