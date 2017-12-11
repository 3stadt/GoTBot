package handlers

import (
	"math/rand"
	"strings"
	"time"

	"github.com/3stadt/GoTBot/src/errors"
)

func (d *deps) Slap() error {
	victim := strings.TrimSpace(d.params)
	if len(d.params) < 1 || strings.ContainsAny(victim, " ") {
		return &fail.TooManyArgs{Max: 1}
	}

	if victim == "himself" || victim == "herself" || victim == "itself" || victim == d.v.Conf["TWITCH_USER"] {
		d.connection.Privmsg(d.channel, "/me slaps "+d.sender+" playfully around with the mighty banhammer...")
		return nil

	}

	rand.Seed(time.Now().Unix())
	objects := []string{
		"a large trout",
		"no visible result",
		"the largest trout ever seen",
		"a barbie doll",
		"a blood stained sack",
		"a chainsaw",
	}
	n := rand.Int() % len(objects)
	d.connection.Privmsg(d.channel, d.sender+" slaps "+victim+" around a bit with "+objects[n]+"!")
	return nil
}
