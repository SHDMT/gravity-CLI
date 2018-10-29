package cmd

import "fmt"

var helpCmd = &Command{
	Use:   "help",
	Short: "find all comminds",
	Run: func(cmd *Command, args []string) {
		fmt.Println("welcome to gravity wallet!")
		fmt.Println("")
		fmt.Println("The following commands will help you implement various" +
			" wallet functions. ")
		fmt.Println("")
		fmt.Println("Usage:    ", "command [arguments]")
		fmt.Println("")
		fmt.Println("The commands are:")
		fmt.Println("")
		for _, cmd := range subCommands {
			fmt.Printf("%7s", "")
			fmt.Printf("%-20s", cmd.Use)
			fmt.Printf("%-20s", cmd.Short)
			fmt.Println("")
		}
		fmt.Println("")
		fmt.Println("Use '[command] -h' for more information about a command.")
		fmt.Println("")

	},
}
