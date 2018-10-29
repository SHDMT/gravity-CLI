package cmd

import (
	"encoding/base64"
	"errors"
	"fmt"
	pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var issuePaymentAccount   string
var issueJson			  string
var issueSend 			  bool

// paymentCmd sends payment to the connected wallet
var issueWithJsonCmd = &Command{
	Use:   "issuecontractwithjson",
	Short: "send issuecontractwithjson command to the connected wallet",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(issueJson) == 0 {
			return errors.New("issueJson must not empty")
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

		//issueJsonBytes,err := hex.DecodeString(issueJson)
		//if err != nil {
		//	fmt.Println("--->issueJson parse filed.")
		//	return
		//}
		issueJson=strings.Replace(issueJson, "\\", "/", -1)
		issueJsonBytes := []byte(issueJson)
		issueWithJsonRequest := &pb.IssueAssetWithJsonRequest{
			PaymentAccount: issuePaymentAccount,
			IssueJson: 		issueJsonBytes,
			Send:			issueSend,
		}

		res, err := c.IssueAssetWithJson(ctx, issueWithJsonRequest)
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(err)
			fmt.Println("<<<<<<<result=======")
			return
		}

		if !issueSend {
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
	issueWithJsonCmd.Flags().StringVarP(&issuePaymentAccount, "issuePaymentAccount", "n", "",
		"account Name string [optional]")
	issueWithJsonCmd.Flags().StringVarP(&issueJson, "issueJson", "j", "",
		"Json-formatted issue message  [optional]")
	issueWithJsonCmd.Flags().BoolVarP(&issueSend, "send", "s", true,
		"send the new unit to network or not [optional]")
}

