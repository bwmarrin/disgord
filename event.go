// Disgord - A scriptable Discord client
// Available at https://github.com/bwmarrin/Disgord

// Copyright 2015 Bruce Marriner <bruce@sqls.net>.  All rights reservec.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains code related to Discord / Disgord events

package disgord

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (dg *Disgord) onEvent(dc *discordgo.Session, st *discordgo.Event) {
	debug(dg.Debug, "%s %s\n", st.Type, st.RawData)
}

func (dg *Disgord) onReady(dc *discordgo.Session, st *discordgo.Ready) {
	debug(dg.Debug, "%+v", st)
}

func (dg *Disgord) onTypingStart(dc *discordgo.Session, st *discordgo.TypingStart) {
	debug(dg.Debug, "%+v", st)
}

func (dg *Disgord) onMessageCreate(dc *discordgo.Session, st *discordgo.Message) {
	debug(dg.Debug, "%+v", st)
	var err error

	// if msg is from self, ignore it entirely.
	if dg.Self == nil || dg.Self.ID == "" {
		dg.Self, err = dg.Discord.User("@me")
		if err != nil {
			fmt.Println("onMessageCreate:: Error fetching self, ", err)
		}
	}
	if st.Author.ID == dg.Self.ID {
		return
	}

	// TODO: Add support for a user/channel ignore list
	// TODO: Add support for a user/channel "listen" list.

	var mentioned bool
	var line = strings.TrimSpace(st.Content)

	// check for @mention
	for _, mention := range st.Mentions {
		if mention.ID == dg.Self.ID {
			mentioned = true

			// strip the mention from the string
			tag := "<@" + dg.Self.ID + ">"
			line = strings.TrimSpace(strings.Replace(line, tag, "", -1))
			break
		}
	}

	// check for prefix mention
	if strings.HasPrefix(line, "-dg ") {
		line = strings.TrimPrefix(line, "-dg ")
		mentioned = true
	}

	// check for private channel "mention"
	_, err = dg.Discord.State.PrivateChannel(st.ChannelID)
	if err == nil {
		mentioned = true
	}

	if mentioned {

		// build a Disgord Message
		m := Message{Message: st}
		m.Content = line

		_, err := dg.Discord.ChannelMessageSend(st.ChannelID, dg.parse(&m))
		if err != nil {
			debug(true, "error sending message: %s", err.Error())
		}
	}

	// process any global stuff that doesn't require a mention
	// like the lua or js stuff below
	// funcs, ok := dg.Hooks["MessageCreate"]
	// for k,v := range funcs
	// v.Callback(dg, dc, st)
	// something, something.
}
