package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"os"
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
	"encoding/json"
	"github.com/SHDMT/gravity/infrastructure/log"
)
var assetForContract string
var listContractsOfAssetCmd = &Command{
	Use:   "listcontracts",
	Short: "List all contracts for asset",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(assetForContract)==0 {
			return errors.New("assetHash has no input")
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

		assetHash,err:= base64Parse(assetForContract)
		if err != nil {
			fmt.Println("Base64 conversion to binary error")
			return
		}
		res, err := c.ListContractsOfAsset(ctx,
			&pb.ListContractsOfAssetRequest{AssetHash:assetHash})
		if err != nil || res==nil{
			fmt.Print("=======result>>>>>>>")
			fmt.Print("Connection to the server failed")
			fmt.Println("<<<<<<<result=======")
			return
		}
		info,err:=json.MarshalIndent(res.ContractWithName, "", " ")
		if err != nil {
			log.Errorf("Json marshal happened error，%v",err)
		}
		fmt.Print("=======result>>>>>>>")
		fmt.Print(string(info))
		fmt.Println("<<<<<<<result=======")
	},
}

func init() {
	listContractsOfAssetCmd.Flags().StringVarP(&assetForContract, "assetHash",
		"a",""," assetHash [required]")
}