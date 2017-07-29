package context

import "github.com/asdine/storm"

var Conf map[string]string
var DB *storm.DB
var PluginDB *storm.DB
var Users storm.Node

const CommandQueueName = "commands"
const UserBucketName = "users"
const DbFile = "gotbot.db"
const PluginDbFile = "gotbotPlugins.db"