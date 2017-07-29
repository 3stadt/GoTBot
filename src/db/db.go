package db

import (
	"github.com/asdine/storm"
	"github.com/3stadt/GoTBot/src/context"
	"github.com/3stadt/GoTBot/src/structs"
	"github.com/imdario/mergo"
	"github.com/3stadt/GoTBot/src/errors"
	"time"
	"strings"
)

func Up() {
	var err error
	context.DB, err = storm.Open(context.DbFile)
	if err != nil {
		panic(err)
	}
	context.PluginDB, err = storm.Open(context.PluginDbFile)
	if err != nil {
		panic(err)
	}
	context.Users = context.DB.From("users")
}

func Down() {
	context.DB.Close()
	context.PluginDB.Close()
}

func UpdateUser(user structs.User) error {
	baseUser := structs.User{}
	err := context.Users.Get(context.UserBucketName, user.Name, &baseUser)
	if err != nil {
		return &fail.NoTargetUser{Name: user.Name}
	}
	mergo.MergeWithOverwrite(&baseUser, user)
	return SetUser(baseUser)
}

func UpdateMessageCount(nick string) error {
	user := structs.User{}
	err := context.Users.Get(context.UserBucketName, nick, &user)
	if err != nil {
		return CreateUser(structs.User{
			Name:         nick,
			MessageCount: 1,
		})
	}
	now := time.Now()
	user.MessageCount++
	user.LastActive = &now
	SetUser(user)
	return nil
}

func CreateUser(user structs.User) error {
	fillDates(&user)
	return SetUser(user)
}

func fillDates(user *structs.User) {
	now := time.Now()
	if user.FirstSeen == nil {
		user.FirstSeen = &now
	}
	if user.LastActive == nil {
		user.LastActive = &now
	}
	if user.LastJoin == nil {
		user.LastJoin = &now
	}
}

func CreateOrUpdateUser(user structs.User) error {
	baseUser := structs.User{}
	err := context.Users.Get(context.UserBucketName, user.Name, &baseUser)
	if err != nil {
		return SetUser(user)
	}
	mergo.MergeWithOverwrite(&baseUser, user)
	return SetUser(baseUser)
}

func GetUser(name string) (*structs.User, error) {
	user := structs.User{}
	err := context.Users.Get(context.UserBucketName, name, &user)
	return &user, err
}

func SetUser(user structs.User) error {
	if strings.TrimSpace(user.Name) == "" {
		return &fail.InvalidStruct{MissingFields: []string{"Name"}}
	}
	if user.MessageCount < 0 {
		user.MessageCount = 0
	}
	err := context.Users.Set(context.UserBucketName, user.Name, user)
	return err
}
