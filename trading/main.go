package main

import (
	"trading/config"
	"trading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
}