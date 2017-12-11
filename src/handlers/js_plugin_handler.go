package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/3stadt/GoTBot/src/db"
	"github.com/3stadt/GoTBot/src/errors"
	"github.com/3stadt/GoTBot/src/res"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
	"github.com/thoj/go-ircevent"
)

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
	vm := otto.New()
	vm.Set("channel", channel)
	vm.Set("sender", sender)
	vm.Set("params", params)

	vm.Set("sendMessage", func(call otto.FunctionCall) otto.Value {
		if len(call.ArgumentList) > 0 {
			msg := call.Argument(0)
			connection.Privmsg(channel, msg.String())
		}
		return otto.Value{}
	})
	vm.Set("getUser", func(call otto.FunctionCall) otto.Value {
		result, _ := vm.ToValue("")
		if len(call.ArgumentList) < 1 {
			return result
		}
		username, err := call.Argument(0).ToString()
		if err != nil {
			return result
		}
		result, _ = vm.ToValue(*getBoltUserAsJson(username, p))
		return result
	})

	vm.Set("setData", func(call otto.FunctionCall) otto.Value {
		result, _ := vm.ToValue("{\"error\": 1}")
		if len(call.ArgumentList) == 2 {
			key := call.Argument(0)
			data := call.Argument(1)
			var dataMap map[string]interface{}
			json.Unmarshal([]byte(data.String()), &dataMap)
			p.PluginDB.Set(bucketName, key, dataMap)
			return result
		}
		failure := fail.NotEnoughArgs{Min: 2}
		result, _ = vm.ToValue(&failure)
		return result
	})
	vm.Set("getData", func(call otto.FunctionCall) otto.Value {
		result, _ := vm.ToValue("{\"error\": 1}")
		if len(call.ArgumentList) == 1 {
			key := call.Argument(0)
			var data map[string]interface{}
			if err := p.PluginDB.Get(bucketName, key, &data); err != nil {
				fmt.Println("Error:")
				fmt.Println(err)
				return result
			}
			var jsonData []byte
			jsonData, err = json.Marshal(data)
			if err != nil {
				return result
			}
			result, _ = vm.ToValue(string(jsonData))
			return result
		}
		return result
	})

	_, err = vm.Run(string(jsData))
	if err != nil {
		fmt.Println("ERROR in javascript file " + filePath + ":")
		fmt.Println(err)
	}
	return nil
}

func getBoltUserAsJson(username string, p *db.Pool) *string {
	emptyJson := "{}"
	userStruct, err := p.GetUser(username)
	if err != nil {
		return &emptyJson
	}
	jUser, err := json.Marshal(*userStruct)
	if err != nil {
		return &emptyJson
	}
	userdata := string(jUser)
	return &userdata
}
