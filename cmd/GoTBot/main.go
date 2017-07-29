package main

import (
	"github.com/thoj/go-ircevent"
	"fmt"
	"strings"
	"github.com/3stadt/GoTBot/src/structs"
	"time"
	"github.com/joho/godotenv"
	"log"
	"github.com/3stadt/GoTBot/src/queue"
	"github.com/3stadt/GoTBot/src/handlers"
	"github.com/3stadt/GoTBot/src/context"
	"strconv"
	"io/ioutil"
	"os"
	"github.com/BurntSushi/toml"
	"github.com/3stadt/GoTBot/src/db"
)

const serverSSL = "irc.chat.twitch.tv:443"

func main() {
	var err error
	_ = initPlugins()

	db.Up()
	defer db.Down()

	context.Conf, err = godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	queue.NewQueue(context.CommandQueueName)
	channel := "#" + context.Conf["TWITCH_CHANNEL"]
	botNick := context.Conf["TWITCH_USER"]
	oauth := context.Conf["OAUTH"]
	debug, debugErr := strconv.ParseBool(context.Conf["DEBUG"])
	if debugErr != nil {
		debug = false
	}
	checkErr(err)
	oauthString := strings.TrimSpace(string(oauth))
	ircConnection := irc.IRC(botNick, botNick)
	ircConnection.VerboseCallbackHandler = debug
	ircConnection.Debug = debug
	ircConnection.UseTLS = true
	ircConnection.Password = oauthString

	ircConnection.AddCallback("001", func(e *irc.Event) {
		ircConnection.SendRaw("CAP REQ :twitch.tv/membership")
		ircConnection.SendRaw("CAP REQ :twitch.tv/commands")
		ircConnection.Join(channel)
		ircConnection.Privmsg(channel, "/mods")
	})

	ircConnection.AddCallback("366", func(e *irc.Event) {})

	ircConnection.AddCallback("NOTICE", func(e *irc.Event) {
		message := e.Message()
		moderatorPrefix := "The moderators of this room are: "
		if strings.HasPrefix(message, moderatorPrefix) {
			context.TwitchMods = strings.Split(strings.TrimPrefix(message, moderatorPrefix), ", ")
		}
	})

	ircConnection.AddCallback("PART", func(e *irc.Event) {
		nick := strings.ToLower(e.Nick)
		if nick == strings.ToLower(botNick) {
			return
		}
		now := time.Now()
		err := db.CreateOrUpdateUser(structs.User{
			Name:     nick,
			LastPart: &now,
		})
		if err != nil {
			panic(err)
		}
	})
	ircConnection.AddCallback("JOIN", func(e *irc.Event) {
		nick := strings.ToLower(e.Nick)
		if nick == strings.ToLower(botNick) {
			return
		}
		now := time.Now()
		err := db.CreateOrUpdateUser(structs.User{
			Name:     nick,
			LastJoin: &now,
		})
		if err != nil {
			panic(err)
		}
	})

	ircConnection.AddCallback("PRIVMSG", func(e *irc.Event) {
		nick := strings.ToLower(e.Nick)
		if err := db.UpdateMessageCount(nick); err != nil {
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
			if _, ok := handlers.CommandMap[command]; ok {
				queue.AddJob(context.CommandQueueName, structs.Job{
					Command: command,
					Channel: channel,
					Sender:  sender,
					Message: message,
					Params:  params,
				})
			} else if _, ok := handlers.PluginCommandMap[command]; ok {
				queue.AddJob(context.CommandQueueName, structs.Job{
					Command: command,
					Channel: channel,
					Sender:  sender,
					Message: message,
					Params:  params,
				})
			}
		}
	})

	go queue.HandleCommand(queue.JobQueue[context.CommandQueueName], ircConnection)

	if err = ircConnection.Connect(serverSSL); err != nil {
		fmt.Printf("Err %s", err)
		return
	}
	ircConnection.Loop()
}
func initPlugins() (error) {
	files, err := ioutil.ReadDir("./custom/plugins")
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			if _, err := os.Stat("./custom/plugins/" + file.Name() + "/config.toml"); !os.IsNotExist(err) {
				tomlData, err := ioutil.ReadFile("./custom/plugins/" + file.Name() + "/config.toml")
				if err != nil {
					log.Fatal(err)
					continue
				}
				var commands structs.Commands
				if _, err := toml.Decode(string(tomlData), &commands); err != nil {
					log.Fatal(err)
				}
				for _, c := range commands.Command {
					if _, ok := handlers.PluginCommandMap[c.Name]; ok {
						handlers.PluginCommandMap[c.Name] = append(handlers.PluginCommandMap[c.Name], "./custom/plugins/"+file.Name()+"/"+c.EntryScript)
					} else {
						handlers.PluginCommandMap[c.Name] = []string{"./custom/plugins/" + file.Name() + "/" + c.EntryScript}
					}
				}
			}
		}
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
