package main

import (
	"trading/app/controllers"
	"trading/config"
	"trading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}