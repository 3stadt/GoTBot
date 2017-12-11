package db

import (
	"strings"
	"time"

	"github.com/3stadt/GoTBot/src/errors"
	"github.com/3stadt/GoTBot/src/structs"
	"github.com/asdine/storm"
	"github.com/imdario/mergo"
)

type Pool struct {
	DB             *storm.DB
	PluginDB       *storm.DB
	Users          *storm.Node
	UserBucketName string
	DbFile         string
	PluginDbFile   string
}

func (p *Pool) Up() (err error) {
	p.UserBucketName = "users"
	p.DB, err = storm.Open(p.DbFile)
	if err != nil {
		return err
	}
	p.PluginDB, err = storm.Open(p.PluginDbFile)
	if err != nil {
		return err
	}
	db := p.DB
	users := db.From("users")
	p.Users = &users
	return nil
}

func (p *Pool) Down() {
	p.DB.Close()
	p.PluginDB.Close()
}

func (p *Pool) UpdateUser(user structs.User) error {
	baseUser := structs.User{}
	userNode := *p.Users
	err := userNode.Get(p.UserBucketName, user.Name, &baseUser)
	if err != nil {
		return &fail.NoTargetUser{Name: user.Name}
	}
	mergo.MergeWithOverwrite(&baseUser, user)
	return p.SetUser(baseUser)
}

func (p *Pool) UpdateMessageCount(nick string) error {
	user := structs.User{}
	userNode := *p.Users
	err := userNode.Get(p.UserBucketName, nick, &user)
	if err != nil {
		return p.CreateUser(structs.User{
			Name:         nick,
			MessageCount: 1,
		})
	}
	now := time.Now()
	user.MessageCount++
	user.LastActive = &now
	p.SetUser(user)
	return nil
}

func (p *Pool) CreateUser(user structs.User) error {
	fillDates(&user)
	return p.SetUser(user)
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

func (p *Pool) CreateOrUpdateUser(user structs.User) error {
	baseUser := structs.User{}
	userNode := *p.Users
	err := userNode.Get(p.UserBucketName, user.Name, &baseUser)
	if err != nil {
		return p.SetUser(user)
	}
	mergo.MergeWithOverwrite(&baseUser, user)
	return p.SetUser(baseUser)
}

func (p *Pool) GetUser(name string) (*structs.User, error) {
	user := structs.User{}
	userNode := *p.Users
	err := userNode.Get(p.UserBucketName, name, &user)
	return &user, err
}

func (p *Pool) SetUser(user structs.User) error {
	if strings.TrimSpace(user.Name) == "" {
		return &fail.InvalidStruct{MissingFields: []string{"Name"}}
	}
	if user.MessageCount < 0 {
		user.MessageCount = 0
	}
	userNode := *p.Users
	err := userNode.Set(p.UserBucketName, user.Name, user)
	return err
}
