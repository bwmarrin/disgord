// Disgord - A scriptable Discord client
// Available at https://github.com/bwmarrin/Disgord

// Copyright 2015 Bruce Marriner <bruce@sqls.net>.  All rights reservec.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains functions and event handlers for Disgord

package disgord

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

// VERSION stores the Version of the Disgord package
const VERSION = "v0.0.0-alpha" // Disgord Version

// Disgord is.. the main disgord package struct
type Disgord struct {
	sync.Mutex

	// Digord Application data
	Debug     bool               // enable Disgord debugging
	Config    config             // Disgord configurable settings
	Commands  map[string]Command // Map of commands
	startTime time.Time          // Time Disgord started

	// Discord Connection
	Discord *discordgo.Session // Discord session
	Self    *discordgo.User    // Authenticated User
}

// Run is the main entry point to Disgord.
// TODO: naming
func (dg *Disgord) Run() {

	// Log startup time
	dg.startTime = time.Now()

	// Set Configuration Defaults
	dg.Config.CommandPrefix = "-dg "

	// Initialize maps
	dg.Commands = make(map[string]Command)
	//dg.Hooks = make(map[string]Hook)

	// Register Built-In Commands
	dg.Commands["help"] = Command{
		Help:        "[<command>]",
		Description: "This help message, or help for a specific command",
		Callback:    help,
	}

	dg.Commands["about"] = Command{
		Help:        "",
		Description: "About this program",
		Callback:    about,
	}

	dg.Commands["stats"] = Command{
		Help:        "",
		Description: "Display program statistics",
		Callback:    stats,
	}

	// TODO: Load JSON based configuration file

	// Parse Run Commands file and execute each Command.
	dg.runCommands()

	// Connect to Discord
	// Has no effect if already connected.
	err := dg.Connect()
	if err != nil {
		fmt.Println("error connecting to discord:", err)
	}

}

// runCommands opens the run commands file and parses all Messages
func (dg *Disgord) runCommands() {

	file, err := os.Open(".dgrc")
	if err != nil {

		if !os.IsNotExist(err) {
			fmt.Println("error opening rcfile: ", err)
		}

		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing rcfile:", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" || strings.HasPrefix(scanner.Text(), "#") {
			continue
		}
		//fmt.Println(dg.Command(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Shutdown will attempt to cleanly shutdown Disgord
func (dg *Disgord) Shutdown() {

	fmt.Println("\n\nStarting Shutdown!")
	err := dg.Discord.Close()
	if err != nil {
		fmt.Println("error closing discord session:", err)
	}
}
