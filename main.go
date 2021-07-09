package main

import (
	"fmt"

	"br.edu.sjc/jarvis/bot"
	"br.edu.sjc/jarvis/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Printf(err.Error())
	}

	bot.Start()

	<-make(chan struct{})
	return
}
