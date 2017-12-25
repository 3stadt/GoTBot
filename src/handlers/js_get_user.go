package handlers

import "github.com/robertkrimen/otto"

func (ircData *ircData) getUser(call otto.FunctionCall) otto.Value {
	result, _ := ircData.vm.ToValue("")
	if len(call.ArgumentList) < 1 {
		return result
	}
	username, err := call.Argument(0).ToString()
	if err != nil {
		return result
	}
	result, _ = ircData.vm.ToValue(*getBoltUserAsJSON(username, ircData.dbPool))
	return result
}
