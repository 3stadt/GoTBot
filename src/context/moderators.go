package context

var TwitchMods []string

func IsTwitchMod(nick string) bool {
	for _, value := range TwitchMods {
		if value == nick {
			return true
		}
	}
	return false
}
