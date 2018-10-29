package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
	"os"
	"encoding/json"
	"github.com/SHDMT/gravity/infrastructure/log"
)
var caddr string
var contractInfoCmd = &Command{
	Use:   "getcontractinfo",
	Short: "Get contract info",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(caddr)==0 {
			return errors.New("caddr has no input")
		}
		return nil
	},
	Run: func(cmd *Command, args []string) {

		c, ctx := newGravityRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()

		caddrHash,err:= base64Parse(caddr)
		if err != nil {
			fmt.Println("Base64 conversion to binary error")
			return
		}
		contractInfoRequest:=&pb.ContractInfoRequest{Caddr:caddrHash}
		res, err := c.GetContractInfo(ctx,contractInfoRequest)

		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print("Connection to the server failed")
			fmt.Println("<<<<<<<result=======")
			return
		}

		info,err:=json.Marshal(res.Contract)
		if err != nil {
			log.Errorf("Json marshal happened error，%v",err)
		}

		fmt.Print("=======result>>>>>>>")
		fmt.Print(string(info))
		fmt.Println("<<<<<<<result=======")
	},
}

func init() {
	contractInfoCmd.Flags().StringVarP(&caddr, "caddr",
		"c",""," caddr [required]")
}