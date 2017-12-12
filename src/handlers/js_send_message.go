package handlers

import (
	"github.com/robertkrimen/otto"
)

func (ircData *ircData) sendMessage(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) > 0 {
		msg := call.Argument(0)
		ircData.c.Privmsg(ircData.channel, msg.String())
	}
	return otto.Value{}
}
