package GoTBot

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/3stadt/GoTBot/src/db"
	"github.com/3stadt/GoTBot/src/handlers"
	"github.com/3stadt/GoTBot/src/queue"
	"github.com/3stadt/GoTBot/src/res"
	"github.com/3stadt/GoTBot/src/structs"
	"github.com/3stadt/GoTBot/src/twitch"
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

const serverSSL = "irc.chat.twitch.tv:443"
const customPluginPath = "./plugins/custom/"
const enginePluginPath = "./plugins/engine/"

func Run() {
	err := initCustomPlugins()
	checkErr(err)
	err = initEnginePlugins()
	checkErr(err)
	p := &db.Pool{
		DbFile:       "gotbot.db",
		PluginDbFile: "gotbotPlugins.db",
	}
	p.Up()
	defer p.Down()
	cfg, err := godotenv.Read()
	rs := &res.Vars{
		Conf:      cfg,
		Constants: res.GetConst(),
	}
	checkErr(err)
	tw, err := connectToTwitch(p, rs)
	checkErr(err)
	tw.Connection.Loop()
}

func connectToTwitch(p *db.Pool, rs *res.Vars) (twitch.Client, error) {
	botNick := rs.Conf["TWITCH_USER"]
	oauth := rs.Conf["OAUTH"]
	debug, debugErr := strconv.ParseBool(rs.Conf["DEBUG"])
	if debugErr != nil {
		debug = false
	}
	tw := twitch.Init(oauth, botNick, rs.Constants.CommandQueueName, debug, p, rs)
	if err := tw.Connection.Connect(serverSSL); err != nil {
		return twitch.Client{}, err
	}
	queue.NewQueue(rs.Constants.CommandQueueName)
	go queue.HandleCommand(queue.JobQueue[rs.Constants.CommandQueueName], tw.Connection, p, rs)
	tw.Connect()
	return tw, nil
}

func initCustomPlugins() error {
	return initPlugins(customPluginPath)
}

func initEnginePlugins() error {
	return initPlugins(enginePluginPath)
}

func initPlugins(folder string) error {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			if _, err := os.Stat(folder + file.Name() + "/config.toml"); !os.IsNotExist(err) {
				tomlData, err := ioutil.ReadFile(folder + file.Name() + "/config.toml")
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
						handlers.PluginCommandMap[c.Name] = append(handlers.PluginCommandMap[c.Name], folder+file.Name()+"/"+c.EntryScript)
					} else {
						handlers.PluginCommandMap[c.Name] = []string{folder + file.Name() + "/" + c.EntryScript}
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
