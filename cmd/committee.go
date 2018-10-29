package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
	"math"
	"encoding/json"
	"github.com/SHDMT/gravity/infrastructure/log"
	"os"
)
var blockeHeight uint32
var committeeCmd = &Command{
	Use:   "getcommittee",
	Short: "get committee by height",
	Args: func(cmd *cobra.Command, args []string) error {
		if blockeHeight==math.MaxUint32 {
			return errors.New("height has no input")
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

		//resHeight, err := c.GetCurrentBlock(ctx, &pb.CurrentBlockRequest{})
		//if resHeight==nil||blockeHeight>resHeight.BlockHeight {
		//	fmt.Print("=======result>>>>>>>")
		//	fmt.Print("The input height is higher than the current maximum" +
		//		" height")
		//	fmt.Println("<<<<<<<result=======")
		//	return
		//}
		res, err := c.GetCommitee(ctx,
			&pb.GetCommiteeRequest{Height:blockeHeight})
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print("Connection to the server failed")
			fmt.Println("<<<<<<<result=======")
			return
		}
		info,err:=json.MarshalIndent(res.Commitee, "", " ")
		if err != nil {
			log.Errorf("Json marshal happened error，%v",err)
		}

		fmt.Print("=======result>>>>>>>")
		fmt.Print(string(info))
		fmt.Println("<<<<<<<result=======")
	},
}

func init() {
	committeeCmd.Flags().Uint32VarP(&blockeHeight, "height",
		"t",math.MaxUint32,"height [required]")
}
