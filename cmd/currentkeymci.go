package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"os"
)
var currentKeyMciCmd = &Command{
	Use:   "getcurrentmci",
	Short: "get current key mci",
	Run: func(cmd *Command, args []string) {

		c, ctx := newGravityRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()

		res, err := c.GetCurrentKeyMci(ctx, &pb.CurrentKeyMciRequest{})
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print("Connection to the server failed")
			fmt.Println("<<<<<<<result=======")
			return
		}

		fmt.Print("=======result>>>>>>>")
		fmt.Print(res.Mci)
		fmt.Println("<<<<<<<result=======")

	},
}
