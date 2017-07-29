package structs

type Command struct {
	Name        string `toml:"name"`
	EntryScript string `toml:"entry_script"`
}

type Commands struct {
	Command []Command
}
