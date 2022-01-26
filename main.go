package main

import (
	"fmt"

	"github.com/OctavianUrsu/go-discord-cryptobot/bot"
	"github.com/OctavianUrsu/go-discord-cryptobot/config"
)

func main() {
	if err := config.ReadConfig(); err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
