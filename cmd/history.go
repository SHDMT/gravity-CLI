package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
	"github.com/spf13/cobra"
	"os"
)

var fromMci uint64
var count uint64

// historyCmd gets history rom the connected wallet
var historyCmd = &Command{
	Use:   "gethistory",
	Short: "gets histories from the connected wallet",
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
		historyRequest := &pb.HistoryRequest{
			AccountName :accountName,
			FromMCI:fromMci,
			Count:count,
		}
		res, err := c.GetPaymentHistory(ctx, historyRequest)
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(err)
			fmt.Println("<<<<<<<result=======")
			return
		}
		fmt.Print("=======result>>>>>>>")
		fmt.Print(string(res.PaymentHistory))
		fmt.Println("<<<<<<<result=======")
	},
}

func init() {
	historyCmd.Flags().StringVarP(&accountName, "account", "a", "",
		"account's name [optional]")
	historyCmd.Flags().Uint64VarP(&fromMci, "mci", "m", 0,
		"select histories from the mci [optional]")
	historyCmd.Flags().Uint64VarP(&count, "count", "c", 0,
		"how many records [optional]")
}
