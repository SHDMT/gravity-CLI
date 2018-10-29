package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"

	"encoding/json"
	"github.com/SHDMT/gravity/infrastructure/log"
	"os"
)

var getPeerInfoCmd = &Command{
	Use:   "getpeerinfo",
	Short: "get connected peers information",
	Run: func(cmd *Command, args []string) {

		c, ctx := newGravityRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()

		getPeerInfoRequest := &pb.GetPeerInfoRequest{}

		res, err := c.GetPeerInfo(ctx, getPeerInfoRequest)
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print("Connection to the server failed")
			fmt.Println("<<<<<<<result=======")
			return
		}
		info,err:=json.MarshalIndent(res.Peers, "", " ")
		if err != nil {
			log.Errorf("Json marshal happened error，%v",err)
		}

		fmt.Print("=======result>>>>>>>")
		fmt.Print(string(info))
		fmt.Println("<<<<<<<result=======")
	},
}
