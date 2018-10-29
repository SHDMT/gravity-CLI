package cmd
import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
	"os"
)
var isGoodUnitString string
var isGoodCmd = &Command{
	Use:   "isgood",
	Short: "Judge unit for good or bad",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(isGoodUnitString)==0 {
			return errors.New("unit has no input")
		}

		//TODO:
		//if len(isGoodUnitString)!=44 {
		//	return errors.New("input unitHash is wrong!")
		//}
		return nil
	},
	Run: func(cmd *Command, args []string) {

		c, ctx := newGravityRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()

		isGoodUnitHash:= removeMarks(isGoodUnitString)

		//todo:判断该unit是否存在
		res, err := c.IsGood(ctx,
			&pb.IsGoodRequest{UnitHash:isGoodUnitHash})
		if err != nil {
			fmt.Println("Connection to the server failed")
			return
		}
		if res.Result {
			fmt.Println("This unit is good unit.")
		}else {
			fmt.Println("This unit is bad unit.")
		}
	},
}

func init() {
	isGoodCmd.Flags().StringVarP(&isGoodUnitString, "unit",
		"u","","unit [required]")
}
