package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/3stadt/GoTBot/src/db"
	"github.com/3stadt/GoTBot/src/res"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
	"github.com/thoj/go-ircevent"
)

type ircData struct {
	c          *irc.Connection
	channel    string
	sender     string
	params     string
	bucketName string
	vm         *otto.Otto
	dbPool     *db.Pool
}

func JsPlugin(filePath string, channel string, sender string, params string, connection *irc.Connection, p *db.Pool, v *res.Vars) error {
	var err error
	var jsData []byte
	var bucketName = filepath.Base(filePath)
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		return err
	}
	if jsData, err = ioutil.ReadFile(filePath); err != nil {
		return err
	}
	ircData := &ircData{
		c:          connection,
		channel:    channel,
		sender:     sender,
		params:     params,
		bucketName: bucketName,
		dbPool:     p,
	}
	vm := otto.New()
	ircData.vm = vm
	vm.Set("channel", channel)
	vm.Set("sender", sender)
	vm.Set("params", params)

	vm.Set("sendMessage", ircData.sendMessage)
	vm.Set("getUser", ircData.getUser)

	vm.Set("setData", ircData.setData)
	vm.Set("getData", ircData.getData)

	_, err = vm.Run(string(jsData))
	if err != nil {
		fmt.Println("ERROR in javascript file " + filePath + ":")
		fmt.Println(err)
	}
	return nil
}

func getBoltUserAsJSON(username string, p *db.Pool) *string {
	emptyJSON := "{}"
	userStruct, err := p.GetUser(username)
	if err != nil {
		return &emptyJSON
	}
	jUser, err := json.Marshal(*userStruct)
	if err != nil {
		return &emptyJSON
	}
	userdata := string(jUser)
	return &userdata
}
