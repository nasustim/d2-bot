package main

import (
	"fmt"
	"os"

	"github.com/nasustim/d2-bot/src/configs"
	"github.com/nasustim/d2-bot/src/repositories/discord"
	"github.com/nasustim/d2-bot/src/repositories/reactions"
)

func main() {
	var d discord.Discord
	var err error
	var stopBot = make(chan bool)

	c, err := configs.Load()
	if err != nil {
		errorHandler(err)
	}
	_, err = d.Auth(c.AuthToken)
	if err != nil {
		errorHandler(err)
	}

	d.AddSimpleReactions(reactions.Reactions())

	err = d.Listen()
	if err != nil {
		errorHandler(err)
	}
	defer d.Close()

	<-stopBot
	return
}

func errorHandler(e error) {
	fmt.Errorf("Error: %s", e)
	os.Exit(1)
}
