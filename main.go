package main

import (
	"discord-cfb-bot/bot"
	"discord-cfb-bot/config"
	"fmt"
)

func main() {
	config.LoadConfig()
	err := bot.Start()
	if err != nil {
		fmt.Println("Error starting the bot:", err)
		return
	}

	select {}
}
