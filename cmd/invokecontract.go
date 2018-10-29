package cmd

import (
"fmt"
pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
"github.com/spf13/cobra"
"encoding/base64"
"os"
	"encoding/hex"
	"strconv"
)

var assetString string
var accountString string
var contractString string
var amountStrs		[]string
var paramString 	string

// paymentCmd sends payment to the connected wallet
var invokeContractCmd = &Command{
	Use:   "invokecontract",
	Short: "send invokecontract command to the connected wallet",
	Args: func(cmd *cobra.Command, args []string) error {
		//if len(amountStrs) == 0 {
		//	fmt.Println(" amount list must not empty")
		//	return nil
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
		asset, err := hex.DecodeString(assetString)
		if err != nil {
			fmt.Println(err)
			return
		}
		contract, err := hex.DecodeString(contractString)
		if err != nil {
			fmt.Println(err)
			return
		}

		amountList := make([]uint64, len(amountStrs))
		for i,amount := range  amountStrs {
			amountI, err := strconv.ParseUint(amount, 10, 64)
			if err != nil {
				fmt.Println(" can't parse amount")
				return
			}
			amountList[i] = amountI
		}

		params,err := hex.DecodeString(paramString)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("param :  %x \n", params)
		invokeContractRequest := &pb.InvokeContractRequest{
			Account:accountString,
			Asset:	asset,
			Contract: contract,
			AmountList:amountList,
			Params:   params,
			Send:		send,
		}

		res, err := c.InvokeContract(ctx, invokeContractRequest)
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(err)
			fmt.Println("<<<<<<<result=======")
			return
		}
		if !send {
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
	invokeContractCmd.Flags().StringVarP(&accountString, "accountName", "n", "",
		"account Name string [required]")
	invokeContractCmd.Flags().StringVarP(&assetString, "asset", "a", "",
		"asset hash string [required]")
	invokeContractCmd.Flags().StringVarP(&contractString, "contractaddress", "c", "",
		"contract address string [required]")
	invokeContractCmd.Flags().StringVarP(&paramString, "params", "p", "",
		"params for the contract [optional]")
	invokeContractCmd.Flags().StringSliceVarP(&amountStrs, "amount list", "m", nil,
		"amount list  [required]")
	invokeContractCmd.Flags().BoolVarP(&send, "send", "s", true,
		"send the new unit to network or not [optional]")
}
