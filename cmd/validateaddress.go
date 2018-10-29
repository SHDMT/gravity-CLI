package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
	"os"
)

var validateAddress string

var validateAddressCmd = &Command{
	Use:   "validateaddress",
	Short: "validate address",
	Args: func(cmd *cobra.Command, args []string) error {
		if validateAddress =="" {
			return errors.New("address has no input")
		}
		return nil
	},
	Run: func(cmd *Command, args []string) {

		c, ctx := newGwalletRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()

		validateAddress = removeMarks(validateAddress)

		res, err := c.ValidateAddress(ctx,
			&pb.ValidateAddressRequest{Address:validateAddress})
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print("Connection to the server failed")
			fmt.Println("<<<<<<<result=======")
			return
		}
		info,err:=json.MarshalIndent(res, "", " ")
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(err)
			fmt.Println("<<<<<<<result=======")
		}
		fmt.Print("=======result>>>>>>>")
		fmt.Print(string(info))
		fmt.Println("<<<<<<<result=======")
	},
}

func init() {
	validateAddressCmd.Flags().StringVarP(&validateAddress, "address", "a", "",
		"address [required]")
}