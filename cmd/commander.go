package cmd

import (
	"fmt"
	"strings"

	"github.com/chzyer/readline"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Command aliases cobra.Command to ease typing
type Command = cobra.Command

// commander serves as handler for command line input
type commander struct {
	rootCmd *Command            // the command at the highest level
	subCmds map[string]*Command // collection of all 2nd-level commands
	rl      *readline.Instance  // instance of Readline
}

// IsReadlineMode checks if the commander runs in ReadLine mode
func (c *commander) IsReadlineMode() bool {
	return nil != c.rl
	//return (nil == c.rootCmd) || !c.rootCmd.HasAvailableSubCommands()
}



// Execute runs the commander in normal mode
func (c *commander) Execute() error {
	if c.IsReadlineMode() {
		c.toggleReadline()
	}

	return c.rootCmd.Execute()
}

// ReadlineAndExecute runs in ReadLine mode
func (c *commander) ReadlineAndExecute() error {
	if !c.IsReadlineMode() {
		c.toggleReadline()
	}

	var line string
	var err error

	//for (nil == err) && c.IsReadlineMode() {
	for c.IsReadlineMode() {
		line, err = c.rl.Readline()
		if nil != err {
			c.rl.Close()
			break
		}

		args := strings.Fields(line)
		if 0 == len(args) {
			continue // skip empty line silently
		}

		hdl, ok := c.subCmds[args[0]]
		if !ok {
			fmt.Println("No such command:", args[0])
			continue
		}

		// trim out the command name
		if len(args) > 1 {
			args = args[1:]
		} else {
			args = nil
		}

		hdl.SetArgs(args)

		err = hdl.Execute()
		resetAllFlagsToDefault(hdl)
	}

	return err
}

// RegisterCommands registers some commands to handle by the commander
func (c *commander) RegisterCommands(commands ...*Command) {
	for _, command := range commands {
		c.subCmds[command.Name()] = command
	}
	c.rootCmd.AddCommand(commands...)
}

// toggleReadline toggles the running mode of the commander
func (c *commander) toggleReadline() {
	var op func(...*Command)

	//if c.rootCmd.HasAvailableSubCommands() {
	if (nil == c.rl) || c.rootCmd.HasAvailableSubCommands() {
		op = c.rootCmd.RemoveCommand

		c.rl, _ = readline.NewEx(&readline.Config{
			Prompt:      AppName + "> ",
			HistoryFile: "readline.history",
			//AutoComplete: pc,
		})
	} else {
		op = c.rootCmd.AddCommand

		if nil != c.rl {
			c.rl.Close()
			c.rl = nil
		}
	}

	for _, v := range c.subCmds {
		op(v)
	}
}


// newDefaultCommander makes a Commander default as non-readline mode
func newDefaultCommander() *commander {
	c := new(commander)
	c.rootCmd = RootCmd
	c.subCmds = make(map[string]*Command)

	// default to non-realine mode
	//c.toggleReadline()

	return c
}

// resetAllFlagsToDefault resets all flags to default values for a given command
func resetAllFlagsToDefault(c *Command) {
	if (nil == c) || !c.HasAvailableFlags() {
		return
	}

	c.Flags().VisitAll(func(f *pflag.Flag) {
		f.Value.Set(f.DefValue)
	})
}
