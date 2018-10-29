package cmd

import (
	"encoding/base64"
	"github.com/spf13/cobra"
	"os"
	"encoding/hex"
	"errors"
	"fmt"

	pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
)

var contractsString 	[]string

// paymentCmd sends payment to the connected wallet
var deployContractCmd = &Command{
	Use:   "deploycontract",
	Short: "send deploycontract command to the connected wallet",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(contractsString) == 0 {
			return errors.New("contract list must not empty")
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

		contracts :=make([][]byte, 0,len(contractsString))
		for _,contractStr := range contractsString {
			contract, err := hex.DecodeString(contractStr)
			if err != nil {
				fmt.Println(err)
				return
			}
			contracts = append(contracts, contract)
		}

		deployRequest := &pb.DeployContractRequest{
			Contracts: contracts,
			Send:		send,
		}

		res, err := c.DeployContract(ctx, deployRequest)
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
	deployContractCmd.Flags().StringVarP(&accountString, "accountName", "n", "",
		"account Name string [optional]")
	deployContractCmd.Flags().StringSliceVarP(&contractsString, "contractsaddress", "c", nil,
		"contracts address hex string [required]")
	deployContractCmd.Flags().BoolVarP(&send, "send", "s", true,
		"send the new unit to network or not [optional]")
}

