package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"os"
)


// paymentCmd sends payment to the connected wallet
var lastKeyUnitCmd = &Command{
	Use:   "getlastkeyunit",
	Short: "get last key unit from dag",
	Run: func(cmd *Command, args []string) {

		c, ctx := newGravityRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()

		lastKeyUnitRequest := &pb.LastKeyUnitRequest{}

		res, err := c.GetLastKeyUnit(ctx, lastKeyUnitRequest)
		if err != nil {
			fmt.Println("Connection to the server failed")
			return
		}
		fmt.Println(res.LastKeyUnitHash)
	},
}


