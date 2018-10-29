package cmd

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
	"strings"
)

var deployPaymentAccount   string
var deployJson			  string
var deploySend 			  bool

// paymentCmd sends payment to the connected wallet
var deployWithJsonCmd = &Command{
	Use:   "deploycontractwithjson",
	Short: "send deploycontractwithjson command to the connected wallet",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(deployJson) == 0 {
			return errors.New("deployJson must not empty")
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

		//deployJsonBytes, err := hex.DecodeString(deployJson)
		//if err != nil {
		//	fmt.Println("--->deployJson parse filed.")
		//	return
		//}
		deployJson=strings.Replace(deployJson, "\\", "/", -1)
		deployJsonBytes := []byte(deployJson)
		deployWithJsonRequest := &pb.DeployContractWithJsonRequest{
			PaymentAccount: deployPaymentAccount,
			DeployJson: 	deployJsonBytes,
			Send:			deploySend,
		}

		res, err := c.DeployContractWithJson(ctx, deployWithJsonRequest)
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(err)
			fmt.Println("<<<<<<<result=======")
			return
		}

		if !deploySend {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(res.Commission)
			fmt.Println("<<<<<<<result=======")
			return
		}
		fmt.Print("=======result>>>>>>>")
		fmt.Print(base64.StdEncoding.EncodeToString(res.UnitHash))
		fmt.Println("<<<<<<<result=======")
	},
}

func init() {
	deployWithJsonCmd.Flags().StringVarP(&deployPaymentAccount, "deployPaymentAccount", "n", "",
		"account Name string [optional]")
	deployWithJsonCmd.Flags().StringVarP(&deployJson, "deployJson", "j", "",
		"Json-formatted deploy message  [optional]")
	deployWithJsonCmd.Flags().BoolVarP(&deploySend, "send", "s", true,
		"send the new unit to network or not [optional]")
}