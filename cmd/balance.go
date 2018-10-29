package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
	"encoding/json"
	"github.com/SHDMT/gravity/infrastructure/log"
	"os"
)

var assetName string
// balanceCmd gets balance rom the connected wallet
var balanceCmd = &Command{
	Use:   "getbalance",
	Short: "gets balance from the connected wallet",
	Run: func(cmd *Command, args []string) {

		c, ctx := newGwalletRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()
		accountName = removeMarks(accountName)
		balanceRequest := &pb.BalanceRequest{
			AccountInfoName:accountName,
			AccountInfoAssetHash:assetName,
		}

		res, err := c.GetBalance(ctx, balanceRequest)
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(err)
			fmt.Println("<<<<<<<result=======")
			return
		}
		info,err:=json.MarshalIndent(res.Balances, "", " ")
		if err != nil {
			log.Errorf("Json marshal happened error，%v",err)
		}
		fmt.Print("=======result>>>>>>>")
		fmt.Print(string(info))
		fmt.Println("<<<<<<<result=======")

	},
}

func init() {
	balanceCmd.Flags().StringVarP(&accountName, "account", "a", "",
		"account's name [optional]")
	balanceCmd.Flags().StringVarP(&assetName, "asset", "e", "",
		"asset's hash [optional]")
}
