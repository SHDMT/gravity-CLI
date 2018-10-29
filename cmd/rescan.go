package cmd

import (
"fmt"
pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
"os"
"github.com/spf13/cobra"
)

var startMCI uint64

// paymentCmd sends payment to the connected wallet
var rescanWalletCmd = &Command{
	Use:   "rescanwallet",
	Short: "rescan wallet address and transactions ",
	Args: func(cmd *cobra.Command, args []string) error {

		return nil
	},
	Run: func(cmd *Command, args []string) {

		c, ctx := newGwalletRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()


		request := &pb.RescanWalletRequest{
			Start:startMCI,
		}

		_, err := c.RescanWallet(ctx, request)
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(err)
			fmt.Println("<<<<<<<result=======")
			return
		}

		fmt.Print("=======result>>>>>>>")
		fmt.Print(true)
		fmt.Println("<<<<<<<result=======")
	},
}

func init() {
	rescanWalletCmd.Flags().Uint64VarP(&startMCI, "start dag mci", "m", 0,
		"rescan started mci [optional]")
}