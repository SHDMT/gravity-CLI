package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"math"
	"os"
)
var miningNum uint32
var startMiningCmd = &Command{
	Use:   "startmining",
	Short: "start mining",
	Args: func(cmd *cobra.Command, args []string) error {
		if miningNum == math.MaxUint32 {
			return errors.New("miningNumber has no input")
		}
		return nil
	},
	Run: func(cmd *Command, args []string) {

		c, ctx := newGravityRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()

		_, err := c.StartMining(ctx,
			&pb.StartMiningRequest{MiningNum:miningNum})
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print("Connection to the server failed")
			fmt.Println("<<<<<<<result=======")
			return
		}

		fmt.Print("=======result>>>>>>>")
		fmt.Print(true)
		fmt.Println("<<<<<<<result=======")

	},
}

func init() {
	startMiningCmd.Flags().Uint32VarP(&miningNum, "miningNumber",
		"n",math.MaxUint32,"miningNumber [required]")
}