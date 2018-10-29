package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
	"github.com/spf13/cobra"
	"os"
	"github.com/pkg/errors"
)

var newAccountName string
var acctType uint32
var accountCmd = &Command{
	Use:   "createnewaccount",
	Short: "Create new  account",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(newAccountName)<1{
			return errors.New("accountName has no input")
		}
		if acctType != 1&&acctType != 0   {
			return errors.New("Wrong account type")
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

		accountName = removeMarks(accountName)
		accountRequest := &pb.CreateNewAccountRequest{
			Account:newAccountName,
			AcctType:acctType,
		}

		//GetNewAddress according account name
		res, err := c.CreateNewAccount(ctx, accountRequest)
		//res, err := c.GetNewAddress(ctx, &pb.AddressRequest{})
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(err)
			fmt.Println("<<<<<<<result=======")
			return
		}

		fmt.Print("=======result>>>>>>>")
		fmt.Print(res.Address)
		fmt.Println("<<<<<<<result=======")
	},
}

func init() {
	accountCmd.Flags().StringVarP(&newAccountName, "account", "a", "",
		"account's name [required]")
	accountCmd.Flags().Uint32VarP(&acctType, "type", "t", 0,
		"account's type [required]")
}
