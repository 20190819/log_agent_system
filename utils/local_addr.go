package utils

import (
	"github.com/yangliang4488/log_agent_system/logger"
	"net"
	"strings"
)

func GetLocalIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		logger.Logger.Error("获取本地 ip 错误: ", err.Error())
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	logger.Logger.Info("本地地址信息: ", localAddr)
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	logger.Logger.Info("本地 ip 地址：",ip)
	return
}
