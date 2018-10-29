package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
	"os"
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
)

var privateKey string

var importCmd = &Command{
	Use:   "importprivatekey",
	Short: "import private key to wallet",
	Args: func(cmd *cobra.Command, args []string) error {
		//TODO:判断privatekey长度或格式
		if privateKey == "" {
			return errors.New("privateKey has no input")
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

		privateKey = removeMarks(privateKey)

		res, err := c.ImportPrivateKey(ctx,
			&pb.ImportPrivateKeyRequest{PrivKey:privateKey,})
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(err)
			fmt.Println("<<<<<<<result=======")
			return
		}
		fmt.Print("=======result>>>>>>>")
		fmt.Print(string(res.Address))
		fmt.Println("<<<<<<<result=======")


	},
}

func init() {
	dumpCmd.Flags().StringVarP(&privateKey, "privateKey", "k", "",
		"privateKey [required]")
}
