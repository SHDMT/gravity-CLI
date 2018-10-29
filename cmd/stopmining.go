package cmd
import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"os"
)
var stopMiningCmd = &Command{
	Use:   "stopmining",
	Short: "stop mining",
	Run: func(cmd *Command, args []string) {

		c, ctx := newGravityRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()

		_, err := c.StopMining(ctx, &pb.StopMiningRequest{})
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print("Connection to the server failed")
			fmt.Println("<<<<<<<result=======")
			return
		}

		fmt.Print("=======result>>>>>>>")
		fmt.Print(true)
		fmt.Println("<<<<<<<result=======")

	},
}