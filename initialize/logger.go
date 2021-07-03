package initialize

import (
	"fmt"
	"time"

	"github.com/childelins/go-gin-api/global"
	"github.com/childelins/go-gin-api/pkg/logger"
)

func InitLogger() error {
	logFile := fmt.Sprintf("storage/logs/app-%s.log", time.Now().Format("2006-01-02"))
	l, err := logger.NewLogger(global.ServerConfig.LogLevel, logFile)
	if err != nil {
		return err
	}

	global.Logger = l
	return nil
}
