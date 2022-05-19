package utils

import (
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func (Util) Logger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			IgnoreRecordNotFoundError: true,         // 忽视
			LogLevel:                  logger.Error, // Log level
			Colorful:                  true,         // 彩色打印
		},
	)
}
