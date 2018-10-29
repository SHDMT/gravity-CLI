package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"os"
)

var progressCmd = &Command{
	Use:   "progress",
	Short: "get synchronization progress",
	Run: func(cmd *Command, args []string) {
		c, ctx := newGravityRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()
		res, err := c.Progress(ctx, &pb.ProgressRequest{})
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(err)
			fmt.Println("<<<<<<<result=======")
			return
		}
		fmt.Print("=======result>>>>>>>")
		fmt.Print(res.ProgressResult)
		fmt.Println("<<<<<<<result=======")

	},
}