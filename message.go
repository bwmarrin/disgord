// Disgord - A scriptable Discord client
// Available at https://github.com/bwmarrin/Disgord

// Copyright 2015 Bruce Marriner <bruce@sqls.net>.  All rights reservec.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains code related to the message
// contents of this file may be merged elsewhere

package disgord

import (
	"github.com/bwmarrin/discordgo"
)

// A Message holds all the information for a Disgord Message that can
// be parsed and routed to commands
type Message struct {
	*discordgo.Message
	// add something for the modified Content
	// with -dg, @user, and what not stripped/cleaned
}
