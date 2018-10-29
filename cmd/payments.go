package cmd

import (
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
	pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
	"fmt"
	"encoding/base64"
	"os"
	"strings"
)

var amounts []uint
var toAddresses []string
var send bool
var flag bool = false


// paymentCmd sends payment to the connected wallet
var paymentsCmd = &Command{
	Use:   "sendpayments",
	Short: "send payments to the connected wallet",
	Args: func(cmd *cobra.Command, args []string) error {

		if amounts ==nil || toAddresses==nil{
			return errors.New("Amounts or addresses has no input")
		}
		if flag {
			toAddresses = toAddresses[1:]
		}
		flag = true
		if len(amounts)!=len(toAddresses) {
			reset()
			return errors.New("Amounts and addresses are not match")
		}
		for _,amount:=range amounts {
			if amount==0 {
				reset()
				return errors.New("payment's amount can't  equal 0")
			}
		}
		for _,toAddress:=range toAddresses{
			if len(toAddress)==0 {
				reset()
				return errors.New("payment's address can't  equal nil")
			}

			//_, _, err := base58.CheckDecode(toAddress)
			//if err != nil {
			//	reset()
			//	return errors.New("receive address is invalid")
			//}
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
		defer reset()
		sendPairLen:=len(amounts)
		sendPairs:=make(map[string]uint64,sendPairLen)
		for i:=0;i<sendPairLen;i++ {

			//toAddress, _, err := base58.CheckDecode(toAddresses[i])
			// if err != nil {
			// 	log.Errorf(" toAddress error : ", err)
			// 	return
			// }
			sendPairs[strings.Replace(toAddresses[i], "\\", "/", -1)]=uint64(amounts[i])
		}
		accountName = removeMarks(accountName)
		SendToManyRequest := &pb.SendToManyRequest{
			SendPaymentAccount : accountName,
			SendPairs:sendPairs,
			Send:send,
		}

		res, err := c.SendToMany(ctx, SendToManyRequest)
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
		fmt.Print(base64.StdEncoding.EncodeToString(res.SendPaymentResult))
		fmt.Println("<<<<<<<result=======")

	},
}

func init() {
	paymentsCmd.Flags().StringVarP(&accountName, "account", "a", "",
		"account's name [optional]")
	paymentsCmd.Flags().UintSliceVarP(&amounts, "amount", "m", nil,
	"payment's amount [required]")
	paymentsCmd.Flags().StringSliceVarP(&toAddresses, "address", "d", nil,
		"send to the address [required]")
	paymentsCmd.Flags().BoolVarP(&send, "send", "s", true,
		"send the new unit to network or not [optional]")
}

func reset()  {
	amounts=nil
	toAddresses=nil
	send = true
}