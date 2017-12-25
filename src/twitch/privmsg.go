package twitch

import (
	"strings"

	"github.com/3stadt/GoTBot/src/handlers"
	"github.com/3stadt/GoTBot/src/queue"
	"github.com/3stadt/GoTBot/src/structs"
	"github.com/thoj/go-ircevent"
)

func (c *Client) Privmsg(e *irc.Event) {
	nick := strings.ToLower(e.Nick)
	if err := c.Pool.UpdateMessageCount(nick); err != nil {
		panic(err)
	}
	message := e.Message()
	if len(message) > 1 && strings.HasPrefix(message, "!") {
		i := strings.Index(message, " ")
		channel := e.Arguments[0]
		sender := nick
		var command string
		var params string

		if i < 0 {
			command = message[1:]
			params = ""
		} else {
			command = message[1:i]
			params = message[i:]
		}
		if err := handlers.Has(command); err == nil {
			queue.AddJob(c.CommandQueueName, structs.Job{
				Command: command,
				Channel: channel,
				Sender:  sender,
				Message: message,
				Params:  params,
			})
		} else if _, ok := handlers.PluginCommandMap[command]; ok {
			queue.AddJob(c.CommandQueueName, structs.Job{
				Command: command,
				Channel: channel,
				Sender:  sender,
				Message: message,
				Params:  params,
			})
		}
	}
}
