// Disgord - A scriptable Discord client
// Available at https://github.com/bwmarrin/Disgord

// Copyright 2015 Bruce Marriner <bruce@sqls.net>.  All rights reservec.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains code related to Disgord configuration

package disgord

// Struct of configurable settings used by Disgord
// NOTE: This will almost surely change format
type config struct {
	OwnerID       string // Discord ID of person running Disgord
	Token         string
	Username      string // Please use Token instead
	Password      string // Please use Token instead
	CommandPrefix string // Command prefix for trigger Disgord parsing
}

// save configuartion to disk
func (c *config) save() {
}

// load configuration from disk
func (c *config) load() {
}

// export configuration to json string
func (c *config) export() {
}

// import configuration from json string
// TODO: naming
func (c *config) importme() {
}
