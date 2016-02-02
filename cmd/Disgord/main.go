// Disgord - A scriptable Discord client
// Available at https://github.com/bwmarrin/Disgord
//
// Copyright 2015 Bruce Marriner <bruce@sqls.net>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/disgord"
)

// Global probram variables
var (
	dg       = disgord.Disgord{}
	Commands = make(map[string]disgord.Command)
)

func init() {

	// Configure Defaults
	dg.Config.CommandPrefix = "-dg "

	// Parse saved configuration

	// Parse from Command Line (overrides Saved Config)
	flag.StringVar(&dg.Config.Token, "t", "", "Discord token.")
	flag.StringVar(&dg.Config.Username, "e", "", "Discord account email.")
	flag.StringVar(&dg.Config.Password, "p", "", "Discord account password.")
	flag.StringVar(&dg.Config.OwnerID, "o", "", "Discord account password.")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Register to track SIGINT/SIGTERM events and listen for them.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	// Start Disgord.
	dg.Run()

	// Register Custom Commands
	for k, v := range Commands {
		dg.Commands[k] = v
	}

	// TODO: Register Custom Hooks

	// Exit on CTRL-C, this should do clean up.
	// TODO: Add check for another method too, like an internal shutdown cmd
	<-sc
	dg.Shutdown()
	fmt.Println("Exiting Normally.")
}
