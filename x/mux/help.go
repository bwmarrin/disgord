package mux

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// Help function provides a build in "help" command that will display a list
// of all registered routes (commands). To use this function it must first be
// registered with the Mux.Route function.
func (m *Mux) Help(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {

	// Set command prefix to display.
	cp := ""
	if ctx.IsPrivate {
		cp = ""
	} else if ctx.HasPrefix {
		cp = m.Prefix
	} else {
		cp = fmt.Sprintf("@%s ", ds.State.User.Username)
	}

	// Sort commands
	maxlen := 0
	keys := make([]string, 0, len(m.Routes))
	cmdmap := make(map[string]*Route)

	for _, v := range m.Routes {

		// Only display commands with a description
		if v.Description == "" {
			continue
		}

		// Calculate the max length of command+args string
		l := len(v.Pattern) // TODO: Add the +args part :)
		if l > maxlen {
			maxlen = l
		}

		cmdmap[v.Pattern] = v

		// help and about are added separately below.
		if v.Pattern == "help" || v.Pattern == "about" {
			continue
		}

		keys = append(keys, v.Pattern)
	}

	sort.Strings(keys)

	// TODO: Learn more link needs to be configurable
	resp := "\n*Commands can be abbreviated and mixed with other text.  Learn more at <https://github.com/bwmarrin/disgord>*\n"
	resp += "```autoit\n"

	v, ok := cmdmap["help"]
	if ok {
		keys = append([]string{v.Pattern}, keys...)
	}

	v, ok = cmdmap["about"]
	if ok {
		keys = append([]string{v.Pattern}, keys...)
	}

	// Add sorted result to help msg
	for _, k := range keys {
		v := cmdmap[k]
		resp += fmt.Sprintf("%s%-"+strconv.Itoa(maxlen)+"s # %s\n", cp, v.Pattern+v.Help, v.Description)
	}

	resp += "```\n"

	ds.ChannelMessageSend(dm.ChannelID, resp)

	return
}
