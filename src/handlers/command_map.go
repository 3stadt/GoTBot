package handlers

import (
	"fmt"

	"github.com/3stadt/GoTBot/src/db"
	"github.com/3stadt/GoTBot/src/errors"
	"github.com/3stadt/GoTBot/src/res"
	"github.com/3stadt/GoTBot/src/structs"
	"github.com/thoj/go-ircevent"
)

type deps struct {
	channel    string
	sender     string
	params     string
	connection *irc.Connection
	p          *db.Pool
	v          *res.Vars
}

var commandMap = map[string]func(*deps) error{
	"goSay":    (*deps).Echo,
	"slap":     (*deps).Slap,
	"stats":    (*deps).Stats,
	"shutdown": (*deps).Stop,
	"stop":     (*deps).Stop,
}

func Has(command string) error {
	if _, ok := commandMap[command]; ok {
		return nil
	}
	return &fail.CommandNotFound{Name: command}
}

func Call(job structs.Job, connection *irc.Connection, p *db.Pool, v *res.Vars) (err error) {
	deps := &deps{
		channel:    job.Channel,
		sender:     job.Sender,
		params:     job.Params,
		connection: connection,
		p:          p,
		v:          v,
	}
	if _, ok := commandMap[job.Command]; ok {
		err = commandMap[job.Command](deps)
		if err != nil {
			fmt.Println(err)
		}
	}
	return &fail.CommandNotFound{Name: job.Command}
}
