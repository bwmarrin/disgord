# dgMux 
[![Discord Gophers](https://img.shields.io/badge/Discord%20Gophers-%23info-blue.svg)](https://discord.gg/0f1SbxBZjYq9jLBk)
<img align="right" src="https://raw.githubusercontent.com/wiki/bwmarrin/disgord/images/gourd.jpg">

dgMux is a simple Discord message route multiplexer that parses a message and 
then executes a matching registered handler. dgMux can be used with 
[Disgord](https://github.com/bwmarrin/disgord) or the 
[DiscordGo](https://github.com/bwmarrin/discordgo) library.

If you would like to help the Disgord or DiscordGo package please use 
[this link](https://discordapp.com/oauth2/authorize?client_id=173113690092994561&scope=bot)
to add the official DiscordGo test bot **dgo** to your server. This provides 
indispensable help to this project.

**For help with this program or general Go discussion, please join the [Discord 
Gophers](https://discord.gg/0f1SbxBZjYq9jLBk) chat server.**

**NOTE** : This is an experimental package and it's likely to have large changes
breaking it's API and compatibility with previous versions.

The goal with dgMux is to create a fairly straight forward and simple 
"command router" that can be added to any DiscordGo bot without much fuss. A 
secondary goal is to keep the route handlers as close to native DiscordGo 
handlers as possible.  So that each example provided can be used by those
just learning the DiscordGo API and easily integrated into their own projects.

Some inspiration was taken from the [chi](https://github.com/go-chi/chi) 
and [httprouter](https://github.com/julienschmidt/httprouter) routers when 
creating dgMux.
