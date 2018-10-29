package cmd

import "fmt"

// exitCmd quits the app
//noinspection ALL
var exitCmd = &Command{
	Use:   "exit",
	Short: "quit the app as you want",
	Run: func(cmd *Command, args []string) {
		if defaultCommander.IsReadlineMode() {
			defaultCommander.toggleReadline()
		}
		fmt.Println("bye bye...")
	},
}
