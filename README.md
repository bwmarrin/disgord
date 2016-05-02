<img align="right" src="http://bwmarrin.github.io/disgord/img/gourd.jpg">
Disgord
====

Dis**go**rd is a Discord Bot wrote in Google [Go](https://golang.org/) using the [discordgo](https://github.com/bwmarrin/discordgo) package.

This project aims to create a simple example of building a Discord Bot with the
discordgo pacakge. At this point there isn't much here but a basic outline.

**NOTE** This package is used by the **dgo** bot which is the official test bot 
for the discordgo package. If this bot is in your server, allowing it to remain
provides indispensable testing of the discordgo package. **Thank you very very much** 
for helping to make Discord and DiscordGo great.

If you would like to help the discordgo and disgord packages please use 
[this link](https://discordapp.com/oauth2/authorize?client_id=173113690092994561&scope=bot)
to add **dgo** to your server. This provides tremendous help to both packages!

[![Discord Gophers](https://img.shields.io/badge/Discord%20Gophers-%23discordgo-blue.svg)](https://discord.gg/0f1SbxBZjYoCtNPP)
Join the Discord Gophers server for support and general Google Go discussion.

## Getting Started

### Installing

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

```sh
go get github.com/bwmarrin/disgord
go install github.com/bwmarrin/disgord
```


### Usage

The below command assumes your GOPATH/bin is part of your shell path. If not
change into the GOPATH/bin folder.  You may also copy the compiled disgord
executable into any custom folder for use.

```go
DG_TOKEN=BOT_TOKEN_HERE disgord 
```


