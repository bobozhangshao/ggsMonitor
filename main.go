package main

import (
	"ggsMonitor/utils/app"
	"ggsMonitor/utils/log"
	"github.com/joho/godotenv"
	"time"
)

func main() {
	app.Error("加载配置文件", godotenv.Load())

	for {
		time.Sleep(1 * time.Second)
		log.Info("干你")
	}
}