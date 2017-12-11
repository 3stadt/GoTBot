package queue

import (
	"fmt"

	"github.com/3stadt/GoTBot/src/db"
	"github.com/3stadt/GoTBot/src/handlers"
	"github.com/3stadt/GoTBot/src/res"
	"github.com/3stadt/GoTBot/src/structs"
	"github.com/thoj/go-ircevent"
)

func commandWorker(job structs.Job, connection *irc.Connection, p *db.Pool, v *res.Vars) {
	var err error
	err = handlers.Call(job, connection, p, v)
	if err != nil {
		err = executeJsFile(job, connection, p, v)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func executeJsFile(job structs.Job, connection *irc.Connection, p *db.Pool, v *res.Vars) error {
	fileNames := handlers.PluginCommandMap[job.Command]
	for _, fileName := range fileNames {
		handlers.JsPlugin(fileName, job.Channel, job.Sender, job.Params, connection, p, v)
	}
	return nil
}

func HandleCommand(c chan structs.Job, connection *irc.Connection, p *db.Pool, v *res.Vars) {
	for {
		commandWorker(<-c, connection, p, v)
	}
}
