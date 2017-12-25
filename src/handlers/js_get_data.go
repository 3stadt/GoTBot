package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/robertkrimen/otto"
)

func (ircData *ircData) getData(call otto.FunctionCall) otto.Value {
	result, _ := ircData.vm.ToValue("{\"error\": 1}")
	if len(call.ArgumentList) == 1 {
		key := call.Argument(0)
		var data map[string]interface{}
		if err := ircData.dbPool.PluginDB.Get(ircData.bucketName, key, &data); err != nil {
			fmt.Println("Error:")
			fmt.Println(err)
			return result
		}
		var jsonData []byte
		jsonData, err := json.Marshal(data)
		if err != nil {
			return result
		}
		result, _ = ircData.vm.ToValue(string(jsonData))
		return result
	}
	return result
}
