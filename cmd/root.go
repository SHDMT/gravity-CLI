package cmd

import "fmt"

// AppName is the name of the app
const AppName = "gravity-cli"

var accountName string

// rootCmd is the entry command name of the app
//noinspection ALL
var RootCmd = &Command{
	Use:   AppName,
	Short: AppName + " is a command line interface to interact with gwallet",
	Run: func(cmd *Command, args []string) {
		fmt.Printf("Hi! Welcome to  %s\n", AppName)
	},
}
