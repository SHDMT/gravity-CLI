# Gravity-CLI  

## Overview  
`gravity-cli` serves as a CLI tool to interact with hubs.

## Command Addition  
The `gravity-cli` program is implemented on the `commander` under `cmd` package. The global `defaultCommander` read in commands and their parameters from the stdin. Then parse the commands and execute the corresponding callback defined in separate files under the `cmd` package, such as `version.go`.

Every command is defined according the specification of the `Command` structure as 
```go
type Command struct {
  Use   string    // the name of the command to execute
  // a short description of the usage for the command, 
  // displayed when call help on the command
  Short string    
  // a long description of the usage for the command
  // displayed when call help on the command
  Long  string    
  // callback for this command, 
  // whose body should be populated with the wanted logic
  Run func(cmd *Command, args []string)   
}
```

So, to add a new command  
1. place your command definition according to `cmd.Command` in a file under the `cmd` package  
2. populate your program logic in the `Run` callback  
3. register your command in the `subCommands` field in the `init.go` file  

Everything is done then!