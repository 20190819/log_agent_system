package tailfile

import (
	"github.com/yangliang4488/log_agent_system/logger"
	"github.com/yangliang4488/log_agent_system/services/etcd"
)

type ManagerClass struct {
	taskMap        map[string]*tailClass
	collectEntries []*etcd.CollectEntry
	newConfChan    <-chan []*etcd.CollectEntry
}

func InitManager(collectEntries []*etcd.CollectEntry, newConfChan <-chan []*etcd.CollectEntry) (err error) {
	manager:= &ManagerClass{
		collectEntries: collectEntries,
		taskMap:        make(map[string]*tailClass, 32),
		newConfChan:    newConfChan,
	}
	for _, collectEntry := range collectEntries {
		// 去重
		if manager.exist(collectEntry.Path) {
			logger.Logger.Warningf("%s 日志已经在收集中", collectEntry.Path)
			continue
		}
		t, err := CreateTailObject(collectEntry.Path, collectEntry.Module, collectEntry.Topic)
		if err != nil {
			logger.Logger.Error("创建 tailClass 失败：", err.Error())
			continue
		}
		// tail 开始执行任务
		go t.run()
		manager.taskMap[collectEntry.Path] = t
	}

	return
}

func (m *ManagerClass) exist(path string) bool {
	for k := range m.taskMap {
		if k == path {
			return true
		}
	}
	return false
}
