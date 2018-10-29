package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"os"
	"encoding/json"
	"github.com/SHDMT/gravity/infrastructure/log"
)
var listAllContractsCmd = &Command{
	Use:   "listallcontractsaddress",
	Short: "list all constracts address and name",
	Run: func(cmd *Command, args []string) {

		c, ctx := newGravityRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()

		res, err := c.ListAllContractsWithName(ctx, &pb.ListAllContractsWithNameRequest{})
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print("Connection to the server failed")
			fmt.Println("<<<<<<<result=======")
			return
		}

		info,err:=json.MarshalIndent(res.AddressAndNames, "", " ")
		if err != nil {
			log.Errorf("Json marshal happened error，%v",err)
		}
		fmt.Print("=======result>>>>>>>")
		fmt.Print(string(info))
		fmt.Println("<<<<<<<result=======")
	},
}