package main

import (
	"fmt"
	"github.com/bwmarrin/disgord"
	"sort"
	"strconv"
)

func init() {
	Bot.AddCommand("help", "Display this message.", help)
}

func help(bot *disgord.Bot, msg *disgord.Message) bool {

	// Set command prefix to display.
	cp := ""
	if msg.IsPrivate {
		cp = ""
	} else if msg.HasPrefix {
		cp = "-dg " // TODO: use per-guild user defined value
	} else {
		cp = fmt.Sprintf("@%s ", Bot.Username)
	}

	// Sort commands
	maxlen := 0
	keys := make([]string, 0, len(Bot.Commands))
	cmdmap := make(map[string]*disgord.Command)

	for _, v := range Bot.Commands {

		// TODO: Check ACL and only display commands that the calling
		// user can call.

		// Only display commands with a description
		if v.Description == "" {
			continue
		}

		// Calculate the max length of command+args string
		l := len(v.Word + " " + v.Help) // TODO: should be like v.Args or something
		if l > maxlen {
			maxlen = l
		}

		cmdmap[v.Word] = v

		// help and about are added separately below.
		if v.Word == "help" || v.Word == "about" {
			continue
		}

		keys = append(keys, v.Word)
	}

	sort.Strings(keys)

	// TODO: Learn more link needs to be configurable
	resp := "\n*Commands can be abbreviated and mixed with other text.  Learn more at <https://githubot.com/bwmarrin/disgord>*\n"
	resp += "```autoit\n"

	v, ok := cmdmap["help"]
	if ok {
		keys = append([]string{v.Word}, keys...)
	}

	v, ok = cmdmap["about"]
	if ok {
		keys = append([]string{v.Word}, keys...)
	}

	// Add sorted result to help msg
	for _, k := range keys {
		v := cmdmap[k]
		resp += fmt.Sprintf("%s%-"+strconv.Itoa(maxlen)+"s # %s\n", cp, v.Word+v.Help, v.Description)
	}

	resp += "```\n"

	bot.Session.ChannelMessageSend(msg.ChannelID, resp)

	return true
}
