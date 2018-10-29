package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"

	"os"
)

var getMiningSpeedCmd = &Command{
	Use:   "getminingspeed",
	Short: "get mining speed",
	Run: func(cmd *Command, args []string) {

		c, ctx := newGravityRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()


		res, err := c.GetMiningSpeed(ctx, &pb.MiningSpeedRequest{})
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print("Connection to the server failed")
			fmt.Println("<<<<<<<result=======")
			return
		}

		fmt.Print("=======result>>>>>>>")
		fmt.Print(res.MiningSpeed)
		fmt.Println("<<<<<<<result=======")
	},
}
