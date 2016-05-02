<img align="right" src="http://bwmarrin.github.io/disgord/img/gourd.jpg">
Disgord
====

Dis**go**rd is a Discord Bot wrote in Google [Go](https://golang.org/) using the [discordgo](https://github.com/bwmarrin/discordgo) package.

This project aims to create a simple example of building a Discord Bot with the
discordgo pacakge. At this point there isn't much here but a basic outline.

Join [![Discord Gophers](https://img.shields.io/badge/Discord%20Gophers-%23discordgo-blue.svg)](https://discord.gg/0f1SbxBZjYoCtNPP)
Discord chat channel for support.

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


