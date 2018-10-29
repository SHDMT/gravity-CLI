package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
	"github.com/spf13/cobra"
	"encoding/base64"
	"os"
)

var unit string

// paymentCmd sends payment to the connected wallet
var sendrawunitCmd = &Command{
	Use:   "sendrawunit",
	Short: "sends rawunit to the connected wallet",
	Args: func(cmd *cobra.Command, args []string) error {
		// if amount == 0 {
		// 	return errors.New("payment's amount can't  equal 0")
		// }
		//if len(toAddressToString)!=64{
		//	return errors.New("address has no input")
		//}
		return nil
	},
	Run: func(cmd *Command, args []string) {

		c, ctx := newGwalletRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()

		// todo 将unit 字符串转为[]byte

		sendrawunitRequest := &pb.SendRawUnitRequest{
			RawUnit:nil,
		}

		res, err := c.SendRawUnit(ctx, sendrawunitRequest)
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(err)
			fmt.Println("<<<<<<<result=======")
			return
		}

		fmt.Println(base64.StdEncoding.EncodeToString(res.SendRawUnitResult))
	},
}

func init() {
	sendrawunitCmd.Flags().StringVarP(&unit, "unit", "u", "",
		"the unit format with hex [required]")
}
