package cmd
import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"encoding/json"
	"github.com/SHDMT/gravity/infrastructure/log"
	"os"
)
var infoCmd = &Command{
	Use:   "getinfo",
	Short: "get dagInfo",
	Run: func(cmd *Command, args []string) {

		c, ctx := newGravityRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()

		res, err := c.GetInfo(ctx, &pb.GetInfoRequest{})
		if err != nil {
			fmt.Println("Connection to the server failed")
			return
		}

		info,err:=json.MarshalIndent(res.DagInfo, "", " ")
		if err != nil {
			log.Errorf("Json marshal happened error，%v",err)
		}
		fmt.Println(string(info))
	},
}

