package handlers

import (
	"strings"
	"github.com/3stadt/GoTBot/src/db"
	"strconv"
	"github.com/thoj/go-ircevent"
)

func Stats(channel string, sender string, params string, connection *irc.Connection) error {
	target := strings.TrimSpace(params)
	if len(params) < 1 || strings.ContainsAny(target, " ") {
		target = sender
	}
	targetUser, err := db.GetUser(target)
	if err != nil {
		return err
	}
	connection.Privmsg(channel, "User " + targetUser.Name + " with " + strconv.Itoa(targetUser.MessageCount) + " messages sent was last active on " + targetUser.LastActive.Format("Mon, Jan 2 15:04:05"))
	return nil
}
