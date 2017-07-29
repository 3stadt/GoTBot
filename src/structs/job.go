package structs

import "github.com/thoj/go-ircevent"

type Job struct {
	Command string
	Channel string
	Sender string
	Message string
	Params string
	Event *irc.Event
}