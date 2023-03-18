package main

import (
	"os"
	"wechatbot/bootstrap"
	"wechatbot/config"

	log "github.com/sirupsen/logrus"
)

func main() {

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	log.SetLevel(log.DebugLevel)
	//log.SetLevel(log.InfoLevel)

	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	log.Info("程序启动")
	err = config.LoadConfig()
	if err != nil {
		log.Warn("没有找到配置文件，尝试读取环境变量")
	}

	wechatEnv := config.GetWechat()
	telegramEnv := config.GetTelegram()
	if wechatEnv != nil && *wechatEnv == "true" {
		bootstrap.StartWebChat()
	} else if telegramEnv != nil {
		bootstrap.StartTelegramBot()
	}

	log.Info("程序退出")
}
