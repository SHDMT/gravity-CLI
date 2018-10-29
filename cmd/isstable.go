package cmd
import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
	"os"
)
var isStableUnitString string
var isStableCmd = &Command{
	Use:   "isstable",
	Short: "Judge unit for stable or not",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(isStableUnitString)==0 {
			return errors.New("unitHash has no input")
		}

		//TODO:
		//if len(isStableUnitString)!=44 {
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

		isStableUnitHash:=removeMarks(isStableUnitString)

		//todo:判断该unit是否存在
		res, err := c.IsStable(ctx,
			&pb.IsStableRequest{UnitHash:isStableUnitHash})
		if err != nil {
			fmt.Println("Connection to the server failed")
			return
		}
		if res.Result {
			fmt.Println("This unit is stable unit.")
		}else {
			fmt.Println("This unit is unstable unit.")
		}
	},
}

func init() {
	isStableCmd.Flags().StringVarP(&isStableUnitString, "unit",
		"u","","unit [required]")
}
