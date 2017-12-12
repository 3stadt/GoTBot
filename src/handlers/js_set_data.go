package handlers

import (
	"encoding/json"

	"github.com/3stadt/GoTBot/src/errors"
	"github.com/robertkrimen/otto"
)

func (ircData *ircData) setData(call otto.FunctionCall) otto.Value {
	result, _ := ircData.vm.ToValue("{\"error\": 1}")
	if len(call.ArgumentList) == 2 {
		key := call.Argument(0)
		data := call.Argument(1)
		var dataMap map[string]interface{}
		json.Unmarshal([]byte(data.String()), &dataMap)
		ircData.dbPool.PluginDB.Set(ircData.bucketName, key, dataMap)
		return result
	}
	failure := fail.NotEnoughArgs{Min: 2}
	result, _ = ircData.vm.ToValue(&failure)
	return result
}
