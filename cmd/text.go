package cmd

import (
	"errors"
	"fmt"
	pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
	"github.com/spf13/cobra"
	"encoding/base64"
	"os"
)

var text string

// textCmd sends text to the connected wallet
var textCmd = &Command{
	Use:   "sendtext",
	Short: "sends text to the connected wallet",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(text) < 1 {
			return errors.New("text has no input")
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

		accountName = removeMarks(accountName)
		sendTextRequest := &pb.SendTextRequest{
			SendTextAccount:accountName,
			TextContent:text,
			Send:send,
		}

		res, err := c.SendText(ctx, sendTextRequest)
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
		fmt.Print(base64.StdEncoding.EncodeToString(res.SendTextResult))
		fmt.Println("<<<<<<<result=======")

	},
}


func init() {
	textCmd.Flags().StringVarP(&accountName, "account", "a",
		"","account's name [optional]")
	textCmd.Flags().StringVarP(&text, "text", "t", "",
		"text for sending [required]")
	textCmd.Flags().BoolVarP(&send, "send", "s", true,
		"send the new unit to network or not [optional]")
}
