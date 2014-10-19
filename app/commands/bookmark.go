package commands

import (
	"fmt"
	"os"

	"../environments"
	"../history_entries"
)

func Bookmark(commandNumber int, bookmarkAs string, env *environments.Environment) {
	if commandNumber < 0 {
		panic("Command number should be >= 0")
	}

	commands, err := history_entries.GetCommands(nil, env)
	if err != nil {
		panic(err)
	}
	if len(commands) <= commandNumber {
		panic("Command number does not exist")
	}

	command := commands[commandNumber-1]
	cmd, _ := command.GetCommand()
	filename := env.GetBookmarkFileName(bookmarkAs)

	file, err := os.Create(filename)
	if err != nil {
		panic(fmt.Sprintf("Cannot create bookmark %s: %v", filename, err))
	}
	defer file.Close()

	file.WriteString(cmd)
}