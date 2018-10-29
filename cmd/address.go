package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
	"github.com/spf13/cobra"
	"os"
)

// AddressCmd gets new  address rom the connected wallet
var addressCmd = &Command{
	Use:   "getnewaddress",
	Aliases:[]string{"NewAddress","getAddress","address"},
	Short: "gets new  address from the connected wallet",
	Args: func(cmd *cobra.Command, args []string) error {
		//加入判断条件
		return nil
	},
	Run: func(cmd *Command, args []string) {

		c, ctx := newGwalletRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()

		accountName = removeMarks(accountName)
		addressRequest := &pb.AddressRequest{
			AccountNameString:accountName,
		}

		//GetNewAddress according account name
		res, err := c.GetNewAddress(ctx, addressRequest)
		//res, err := c.GetNewAddress(ctx, &pb.AddressRequest{})
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(err)
			fmt.Println("<<<<<<<result=======")
			return
		}
		fmt.Print("=======result>>>>>>>")
		fmt.Print(res.AddressString)
		fmt.Println("<<<<<<<result=======")
	},
}

func init() {
	addressCmd.Flags().StringVarP(&accountName, "account", "a", "",
		"account's name [optional]")
}
