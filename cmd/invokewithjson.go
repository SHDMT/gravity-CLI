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

var invokePaymentAccount   string
var invokeJson			  string
var invokeAmount 		  uint64
var invokeSend 			  bool

// paymentCmd sends payment to the connected wallet
var invokeWithJsonCmd = &Command{
	Use:   "invokecontractwithjson",
	Short: "send invokecontractwithjson command to the connected wallet",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(invokeJson) == 0 {
			return errors.New("invokeJson list must not empty")
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

		//invokeJsonBytes, err := hex.DecodeString(invokeJson)
		//if err != nil {
		//	fmt.Println("--->invokeJson parse filed.")
		//	return
		//}
		invokeJson=strings.Replace(invokeJson, "\\", "/", -1)
		invokeJsonBytes := []byte(invokeJson)
		invokeWithJsonRequest := &pb.InvokeContractWithJsonRequest{
			PaymentAccount: invokePaymentAccount,
			InvokeJson: 	invokeJsonBytes,
			Amount:			invokeAmount,
			Send:			invokeSend,
		}

		res, err := c.InvokeContractWithJson(ctx, invokeWithJsonRequest)
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(err)
			fmt.Println("<<<<<<<result=======")
			return
		}

		if !invokeSend{
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
	invokeWithJsonCmd.Flags().StringVarP(&invokePaymentAccount, "invokePaymentAccount", "n", "",
		"account Name string [optional]")
	invokeWithJsonCmd.Flags().StringVarP(&invokeJson, "invokeJson", "j", "",
		"Json-formatted invoke message  [required]")
	invokeWithJsonCmd.Flags().Uint64VarP(&invokeAmount, "invokeAmount", "a", 0,
		"payment amount of invoke message  [optional]")
	invokeWithJsonCmd.Flags().BoolVarP(&invokeSend, "send", "s", true,
		"send the new unit to network or not [optional]")
}
