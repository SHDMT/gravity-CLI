package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
	"encoding/json"
	"github.com/SHDMT/gravity/infrastructure/log"
	"math"
	"bytes"
	"encoding/gob"
	"os"
	"github.com/SHDMT/gravity/platform/gpow/commonstructs"
)
var height uint32
var blockCmd = &Command{
	Use:   "getblock",
	Short: "get block by height",
	Args: func(cmd *cobra.Command, args []string) error {
		if height==math.MaxUint32 {
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


		resHeight, err := c.GetCurrentBlock(ctx,&pb.CurrentBlockRequest{})
		if resHeight.BlockHeight<height {
			fmt.Println("The input height is higher than the current maximum" +
				" height")
			return
		}
		res, err := c.GetBlockWithCert(ctx,
			&pb.BlockWithCertRequest{Height:height})
		if err != nil {
			fmt.Println("Connection to the server failed")
			return
		}

		block := new(commonstructs.Block)
		buf := bytes.NewBuffer(res.Block)
		gob.NewDecoder(buf).Decode(block)

		info,err:=json.MarshalIndent(block, "", " ")
		if err != nil {
			log.Errorf("Json marshal happened error，%v",err)
		}

		fmt.Print("=======result>>>>>>>")
		fmt.Print(string(info))
		fmt.Println("<<<<<<<result=======")
	},
}

func init() {
	blockCmd.Flags().Uint32VarP(&height, "height",
		"t",math.MaxUint32,"height [required]")
}
