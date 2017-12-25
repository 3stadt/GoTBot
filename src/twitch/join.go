package twitch

import (
	"strings"
	"time"

	"github.com/3stadt/GoTBot/src/structs"
	"github.com/thoj/go-ircevent"
)

func (c *Client) Join(e *irc.Event) {
	nick := strings.ToLower(e.Nick)
	if nick == strings.ToLower(c.Nick) {
		return
	}
	now := time.Now()
	err := c.Pool.CreateOrUpdateUser(structs.User{
		Name:     nick,
		LastJoin: &now,
	})
	if err != nil {
		panic(err)
	}
}
