package queue

import (
	"github.com/3stadt/GoTBot/src/structs"
	"github.com/3stadt/GoTBot/src/handlers"
	"github.com/thoj/go-ircevent"
	"fmt"
)

func commandWorker(job structs.Job, connection *irc.Connection) {
	var msg *structs.Message
	var err error
	if _, ok := handlers.CommandMap[job.Command]; ok {
		err = handlers.CommandMap[job.Command](job.Channel, job.Sender, job.Params, connection)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		if err := executeJsFile(job, connection); err != nil {
			fmt.Println(err)
		}
	}
	if msg != nil {
		connection.Privmsg(msg.Channel, msg.Message)
	}
}
func executeJsFile(job structs.Job, connection *irc.Connection) (error) {
	fileNames := handlers.PluginCommandMap[job.Command]
	for _, fileName := range fileNames {
		handlers.JsPluginHandler(fileName, job.Channel, job.Sender, job.Params, connection)
	}
	return nil
}

func HandleCommand(c chan structs.Job, connection *irc.Connection) {
	for {
		commandWorker(<-c, connection)
	}
}
