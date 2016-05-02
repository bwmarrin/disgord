package main

/* Disgord Spew Plugin
 * This plugin demonstrates attaching to the OnEvent handler and printing
 * out all discord events using the spew package.
 */

import (
	"github.com/bwmarrin/discordgo"
	"github.com/davecgh/go-spew/spew"
)

func init() {
	Bot.AddHandler(onEvent)
}

func onEvent(s *discordgo.Session, e *discordgo.Event) {
	spew.Dump(e)
}
