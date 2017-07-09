package handlers

import (
	"strings"
	"time"
	"math/rand"
	"github.com/3stadt/GoTBot/globals"
	"github.com/3stadt/GoTBot/structs"
	"github.com/3stadt/GoTBot/errors"
)

func Slap(channel string, sender string, params string) (*structs.Message, error) {
	victim := strings.TrimSpace(params)
	if len(params) < 1 || strings.ContainsAny(victim, " ") {
		return nil, &fail.TooManyArgs{Max: 1}
	}

	if victim == "himself" || victim == "herself" || victim == "itself" || victim == globals.Conf["TWITCH_USER"] {
		return &structs.Message{
			Channel: channel,
			Message: "/me slaps " + sender + " playfully around with the mighty banhammer...",
		}, nil

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
	return &structs.Message{
		Channel: channel,
		Message: sender + " slaps " + victim + " around a bit with " + objects[n] + "!",
	}, nil
}
