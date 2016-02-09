// Disgord - A scriptable Discord client
// Available at https://github.com/bwmarrin/Disgord

// Copyright 2015 Bruce Marriner <bruce@sqls.net>.  All rights reservec.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains code related to the primary command handler

package disgord

import (
	"fmt"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

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

	return "Unknown command, try help"
}

func help(dg *Disgord, m *Message) (msg string) {

	msg = "Disgord Help:\n\n" // TODO: Allow Disgord to be a configurable botname

	if len(dg.Commands) < 1 {
		return
	}

	maxlen := 0

	// Sort commands
	keys := make([]string, 0, len(dg.Commands))
	for k, v := range dg.Commands {
		// TODO: Check ACL and only display commands that the calling
		// user can call.

		l := len(k + " " + v.Help)
		if l > maxlen {
			maxlen = l
		}

		keys = append(keys, k)
	}
	sort.Strings(keys)

	// TODO: can we push about/help to the top?

	// Add sorted result to help msg
	for _, k := range keys {
		v := dg.Commands[k]
		msg += fmt.Sprintf("```%-"+strconv.Itoa(maxlen)+"s : %s```", k+" "+v.Help, v.Description)
	}

	return
}

func about(dg *Disgord, m *Message) string {

	return "\n" +
		"Disgord " + VERSION + " (github.com/bwmarrin/Disgord)\n" +
		"Owner is <@" + dg.Config.OwnerID + ">\n"
}

func stats(dg *Disgord, m *Message) (msg string) {

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	msg = "```\n" +
		fmt.Sprintf("Disgord      : %s\n", VERSION) +
		fmt.Sprintf("DiscordGo    : v%s\n", discordgo.VERSION) +
		fmt.Sprintf("Uptime       : %s\n", time.Now().Sub(dg.startTime)) +
		fmt.Sprintf("Processes    : %d\n", runtime.NumGoroutine()) +
		fmt.Sprintf("HeapAlloc    : %.2fMB\n", float64(mem.HeapAlloc)/1048576) +
		fmt.Sprintf("Total Sys    : %.2fMB\n", float64(mem.Sys)/1048576)

	if dg.Discord.StateEnabled && dg.Discord.State != nil {
		guilds := len(dg.Discord.State.Guilds)
		channels := 0
		members := 0
		for _, v := range dg.Discord.State.Guilds {
			channels += len(v.Channels)
			members += len(v.Members)
		}

		msg += fmt.Sprintf("Guilds       : %d\n", guilds)
		msg += fmt.Sprintf("Channels     : %d\n", channels)
		msg += fmt.Sprintf("Members      : %d\n", members)
	}

	msg += "```"

	return
}
