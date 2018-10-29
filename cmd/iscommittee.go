package cmd
import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"os"
)

var isCommitteeCmd = &Command{
	Use:   "iscommittee",
	Short: "Judge is or not a committee member",
	Run: func(cmd *Command, args []string) {

		c, ctx := newGravityRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()

		res, err := c.IsCommittee(ctx, &pb.IsCommitteeRequest{})
		if err != nil {
			fmt.Println("Connection to the server failed")
			return
		}
		fmt.Print("=======result>>>>>>>")
		fmt.Print(res.Result)
		fmt.Println("<<<<<<<result=======")
	},
}

