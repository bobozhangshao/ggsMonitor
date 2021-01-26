package main

import (
	"ggsMonitor/app"
	"ggsMonitor/app/log"
	"github.com/joho/godotenv"
	"time"
)

func main() {
	app.Error("配置文件", godotenv.Load())

	for {
		time.Sleep(1 * time.Second)
		log.Info("干你")
	}
}