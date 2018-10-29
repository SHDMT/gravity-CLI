package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
	"os"
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
)

var dumpAddress string

var dumpCmd = &Command{
	Use:   "exportprivatekey",
	Short: "export private key from wallet",
	Args: func(cmd *cobra.Command, args []string) error {
		//TODO:判断address长度或格式
		if dumpAddress == "" {
			return errors.New("address has no input")
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

		dumpAddress = removeMarks(dumpAddress)

		res, err := c.DumpPrivateKey(ctx, &pb.DumpPrivateKeyRequest{Address:dumpAddress,})
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(err)
			fmt.Println("<<<<<<<result=======")
			return
		}
		fmt.Print("=======result>>>>>>>")
		fmt.Print(res.PrivKey)
		fmt.Println("<<<<<<<result=======")

	},
}

func init() {
	dumpCmd.Flags().StringVarP(&dumpAddress, "address", "a", "",
		"address [required]")
}
