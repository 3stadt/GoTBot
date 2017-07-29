package handlers

import (
	"github.com/thoj/go-ircevent"
)

var CommandMap = map[string]func(channel string, sender string, params string, connection *irc.Connection) error {
	"goSay": Echo,
	"slap":  Slap,
	"stats": Stats,
	"shutdown": Stop,
	"stop": Stop,
}