package twitch

import (
	"strings"

	"github.com/3stadt/GoTBot/src/db"
	"github.com/3stadt/GoTBot/src/res"
	"github.com/thoj/go-ircevent"
)

type Client struct {
	Connection       *irc.Connection
	Nick             string
	Moderators       []string
	CommandQueueName string
	Vars             *res.Vars
	Pool             *db.Pool
}

func Init(oauth string, nick string, commandQueueName string, debug bool, p *db.Pool, res *res.Vars) Client {
	client := Client{
		Connection: irc.IRC(nick, nick),
		Nick:       nick,
		Vars:       res,
		Pool:       p,
	}
	client.Connection.Password = oauth
	client.CommandQueueName = commandQueueName
	client.Moderators = []string{}
	oauthString := strings.TrimSpace(string(oauth))
	client.Connection.VerboseCallbackHandler = debug
	client.Connection.Debug = debug
	client.Connection.UseTLS = true
	client.Connection.Password = oauthString
	client.Connection.AddCallback("001", client.Connected)
	client.Connection.AddCallback("366", func(e *irc.Event) {})
	client.Connection.AddCallback("NOTICE", client.Notice)
	client.Connection.AddCallback("PART", client.Part)
	client.Connection.AddCallback("JOIN", client.Join)
	client.Connection.AddCallback("PRIVMSG", client.Privmsg)
	return client
}

func (c *Client) Connect() {
	c.Connection.Loop()
}

func (c *Client) Disconnect() {
	c.Connection.Quit()
}
