package handlers

import (
	"strings"
	"time"
	"math/rand"
	"github.com/3stadt/GoTBot/src/errors"
	"github.com/thoj/go-ircevent"
	"github.com/3stadt/GoTBot/src/db"
	"github.com/3stadt/GoTBot/src/res"
)

func Slap(channel string, sender string, params string, connection *irc.Connection, p *db.Pool, v *res.Vars) error {
	victim := strings.TrimSpace(params)
	if len(params) < 1 || strings.ContainsAny(victim, " ") {
		return &fail.TooManyArgs{Max: 1}
	}

	if victim == "himself" || victim == "herself" || victim == "itself" || victim == v.Conf["TWITCH_USER"] {
		connection.Privmsg(channel, "/me slaps "+sender+" playfully around with the mighty banhammer...")
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
	connection.Privmsg(channel, sender+" slaps "+victim+" around a bit with "+objects[n]+"!")
	return nil
}
