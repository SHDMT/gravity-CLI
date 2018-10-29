package main

import (
	"fmt"
	"github.com/SHDMT/gravity-cli/cmd"
)

func main() {
	cmd.Execute()
	if err := cmd.ReadlineAndExecute(); nil != err {
		fmt.Println(err)
	}
}

//func init(){
//	log.Init(gconfig.Parameters.MaxLogSize, gconfig.Parameters.LogLevel,
//		gconfig.DefaultLogDir, log.Stdout)
//}
