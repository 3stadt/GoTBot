package handlers

import (
	"strconv"
	"strings"
)

func (d *deps) Stats() error {
	target := strings.TrimSpace(d.params)
	if len(d.params) < 1 || strings.ContainsAny(target, " ") {
		target = d.sender
	}
	targetUser, err := d.p.GetUser(target)
	if err != nil {
		return err
	}
	d.connection.Privmsg(d.channel, "User "+targetUser.Name+" with "+strconv.Itoa(targetUser.MessageCount)+" messages sent was last active on "+targetUser.LastActive.Format("Mon, Jan 2 15:04:05"))
	return nil
}
