package cmd

import (
	"errors"
	"fmt"
	pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
	"github.com/spf13/cobra"
	"math"
	"encoding/base64"
	"os"
)

var unitHashToString string
var messageId uint32

// messageCmd gets info of message from the connected wallet
var messageCmd = &Command{
	Use:   "getmessageinfo",
	Short: "gets info of message from the connected wallet",
	Args: func(cmd *cobra.Command, args []string) error {
		//if len(unitHashToString) > 0 {
		//	return errors.New("There was an error in the input of unit")
		//}
		if messageId == math.MaxUint32 {
			return errors.New("message's id has no input")
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

		unitHash,err := base64Parse(unitHashToString)
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(err)
			fmt.Println("<<<<<<<result=======")
			return
		}
		messageInfoRequest := &pb.MessageInfoRequest{
			UnitHash:unitHash,
			MessageId:messageId,
		}

		res, err := c.GetPaymentMessageInfo(ctx, messageInfoRequest)
		if err != nil {
			fmt.Println("Connection to the server failed")
			return
		}
		fmt.Print("=======result>>>>>>>")
		fmt.Print(base64.StdEncoding.EncodeToString(res.PaymentInfo))
		fmt.Println("<<<<<<<result=======")
	},
}

func init() {
	messageCmd.Flags().StringVarP(&unitHashToString, "unit", "u", "",
		"unit's tag [required]")
	//math.MaxUint32==4294967295
	messageCmd.Flags().Uint32VarP(&messageId, "id", "i", math.MaxUint32,
		"message's id [required]")
}
