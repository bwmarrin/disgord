// Disgord - A scriptable Discord client
// Available at https://github.com/bwmarrin/Disgord

// Copyright 2015 Bruce Marriner <bruce@sqls.net>.  All rights reservec.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains code related to the primary command handler

package disgord

import "strings"

// A Command struct holds information about an individual command
// or a domain of commands (see Command.Commands)
type Command struct {
	Help        string // help message added after command
	Description string // Description of command
	ExtraHelp   string // Detailed help message

	// Map of sub-commands
	Commands map[string]Command

	// Callback function that will parse/handle this command
	// TODO: Review name.
	Callback func(*Disgord, *Message) string

	// TODO: Add ACL stuff
	// Need ::
	// List of allowed/blocked users
	// List of allowed/blocked roles
	// List of allowed/blocked guilds
	// List of allowed/blocked channels
	// Flag saying only "Owner" can run this
}

// parse a mesage and route it to the appropriate command handler
func (dg *Disgord) parse(msg *Message) string {

	split := strings.SplitN(msg.Content, " ", 2)
	cmd := split[0]

	arg := ""

	if len(split) == 2 {
		arg = strings.TrimSpace(split[1])
	}

	c, ok := dg.Commands[cmd]
	if ok {
		// TODO: Check ACL

		if c.Commands != nil {
			// TODO: Check for sub-Commands
		}

		msg.Content = arg
		return c.Callback(dg, msg)
	}

	if dg.DefaultCommand != nil {
		return dg.DefaultCommand.Callback
	}
}
