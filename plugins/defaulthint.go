package main

import (
	"github.com/bwmarrin/disgord"
)

func init() {
	cmd, _ := Bot.AddCommand("hint", "A little hint on using this Bot.", hint)
	Bot.DefaultCommand = cmd
}

func hint(bot *disgord.Bot, msg *disgord.Message) bool {

	var resp string

	if msg.IsPrivate {
		resp = "Try entering **help** to see a list of commands this bot supports."
	} else {
		resp = "Try entering **@" + bot.Username + " help** to see a list of commands this bot supports."
	}

	bot.Session.ChannelMessageSend(msg.ChannelID, resp)

	return true
}
