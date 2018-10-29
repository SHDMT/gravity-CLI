package cmd

import "github.com/chzyer/readline"

var clearCmd = &Command{
	Use:   "clear",
	Short: "clear up the screen",
	Run: func(cmd *Command, args []string) {
		if (nil != defaultCommander) && (nil != defaultCommander.rl) {
			readline.ClearScreen(defaultCommander.rl.Stdout())
		}
	},
}
