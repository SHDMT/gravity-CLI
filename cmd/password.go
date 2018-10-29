package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
	"os"
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
)

var oldPassword string
var newPassword string

// paymentCmd sends payment to the connected wallet
var updatePasswordCmd = &Command{
	Use:   "updatepassword",
	Short: "update password for your wallet",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(oldPassword) == 0 {
			return errors.New("new password has no input")
		}
		if len(newPassword) == 0 {
			return errors.New("new password has no input")
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

		oldPassword = removeMarks(oldPassword)
		newPassword = removeMarks(newPassword)
		updatePasswordRequest := &pb.UpdatePasswordRequest{
			OldPassword:oldPassword,
			NewPassword:newPassword,
		}

		_, err := c.UpdatePassword(ctx, updatePasswordRequest)
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
	updatePasswordCmd.Flags().StringVarP(&oldPassword, "oldpassword", "o", "",
		"new password [required]")
	updatePasswordCmd.Flags().StringVarP(&newPassword, "newpassword", "n", "",
		"new password [required]")
}

