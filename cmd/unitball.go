package cmd
import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
	"os"
)
var unitString string
var unitBallCmd = &Command{
	Use:   "getunitball",
	Short: "get ball by unit",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(unitString)==0 {
			return errors.New("unitHash has no input")
		}

		//TODO:
		//if len(unitString)!=44 {
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

		unitHash:=removeMarks(unitString)
		res, err := c.GetUnitBall(ctx,
			&pb.GetUnitBallRequest{UnitHash:unitHash})
		if err != nil {
			fmt.Println("Connection to the server failed")
			return
		}
		fmt.Println(res.BallHash)
	},
}

func init() {
	unitBallCmd.Flags().StringVarP(&unitString, "unit",
		"u","","unit [required]")
}
