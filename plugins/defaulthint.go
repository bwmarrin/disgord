package main

import (
	"github.com/bwmarrin/disgord/bot"
)

func init() {
	cmd, _ := Bot.AddCommand("hint", "A little hint on using this Bot.", hint)
	Bot.DefaultCommand = cmd
}

func hint(b *bot.Bot, m *bot.Message) bool {

	var resp string

	if m.IsPrivate {
		resp = "Try entering **help** to see a list of commands this bot supports."
	} else {
		resp = "Try entering **@" + b.Name + " help** to see a list of commands this bot supports."
	}

	b.ChannelMessageSend(m.ChannelID, resp)

	return true
}
