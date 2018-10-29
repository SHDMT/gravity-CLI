package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
	"os"
	"encoding/json"
	"github.com/SHDMT/gravity/infrastructure/log"
)

var allAssetsCmd = &Command{
	Use:   "getallassets",
	Short: "Get all assets",
	Run: func(cmd *Command, args []string) {

		c, ctx := newGwalletRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()

		res, err := c.GetAllAssets(ctx, &pb.AllAssetsRequest{})
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(err)
			fmt.Println("<<<<<<<result=======")
			return
		}
		info,err:=json.MarshalIndent(res.Assets, "", " ")
		if err != nil {
			log.Errorf("Json marshal happened error，%v",err)
		}
		fmt.Print("=======result>>>>>>>")
		fmt.Print(string(info))
		fmt.Println("<<<<<<<result=======")
	},
}
