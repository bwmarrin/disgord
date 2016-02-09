// Disgord - A scriptable Discord client
// Available at https://github.com/bwmarrin/Disgord

// Copyright 2015 Bruce Marriner <bruce@sqls.net>.  All rights reservec.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains functions and event handlers for Disgord

package disgord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Connect connects to Discord.
func (dg *Disgord) Connect() (err error) {

	dg.Lock()
	defer dg.Unlock()

	// If Discord is not nil, assume we're already
	// connected.  Might need to make this smarter.
	if dg.Discord != nil {
		return
	}

	dg.Discord, err = discordgo.New(dg.Config.Username, dg.Config.Password, dg.Config.Token)
	if err != nil {
		fmt.Println("error connecting to discord: ", err)
		return
	}

	// Save token for future use.
	dg.Discord.Token = dg.Config.Token

	// Get Authenticated User's information
	dg.Self, err = dg.Discord.User("@me")
	if err != nil {
		fmt.Println("error fetching self, ", err)
		return
	}

	// Register event handlers
	dg.Discord.OnEvent = dg.onEvent
	dg.Discord.OnReady = dg.onReady
	dg.Discord.OnTypingStart = dg.onTypingStart
	dg.Discord.OnMessageCreate = dg.onMessageCreate

	// Open websocket connection
	err = dg.Discord.Open()
	if err != nil {
		fmt.Println("Discord Error: " + err.Error())
		return
	}

	return
}
