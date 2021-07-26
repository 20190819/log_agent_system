package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	// 设置日志输出为os.Stdout
	Logger.Out = os.Stdout
	Logger.Level = logrus.DebugLevel

	Logger.Info("init logger success")
}
