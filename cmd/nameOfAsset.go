package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"os"
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
)
var assetForName string
var getNameOfAssetCmd = &Command{
	Use:   "getnameofasset",
	Short: "A simple introduction",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(assetForName)==0 {
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

		assetHash,err:= base64Parse(assetForName)
		if err != nil {
			fmt.Println("Base64 conversion to binary error")
			return
		}
		res, err := c.GetNameOfAsset(ctx,&pb.GetNameOfAssetRequest{AssetHash:assetHash})
		if err != nil || res==nil{
			fmt.Print("=======result>>>>>>>")
			fmt.Print("Connection to the server failed")
			fmt.Println("<<<<<<<result=======")
			return
		}
		if len(res.Err)!=0 {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(res.Err)
			fmt.Println("<<<<<<<result=======")
			return
		}

		fmt.Print("=======result>>>>>>>")
		fmt.Print(res.Name)
		fmt.Println("<<<<<<<result=======")
	},
}

func init() {
	getNameOfAssetCmd.Flags().StringVarP(&assetForName, "assetHash",
		"a",""," assetHash [required]")
}