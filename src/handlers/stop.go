package handlers

import (
	"github.com/thoj/go-ircevent"
	"errors"
	"github.com/3stadt/GoTBot/src/context"
	"strings"
)

func Stop(channel string, sender string, params string, connection *irc.Connection) error {
	if sender == strings.TrimPrefix(channel, "#") || context.IsTwitchMod(sender) {
		connection.Privmsg(channel, "Shutting down bot...")
		connection.Quit()
		return nil
	}
	return errors.New("Insufficient permissions")
}
