package cmd
import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
	"encoding/json"
	"github.com/SHDMT/gravity/infrastructure/log"
)
var ballHashString string
var ballCmd = &Command{
	Use:   "getball",
	Aliases:[]string{"get"},
	Short: "get ball by ballHash",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(ballHashString)==0 {
			return errors.New("ballHash has no input")
		}

		//TODO:
		//if len(ballHashString)!=44 {
		//	return errors.New("input ballHash is wrong!")
		//}
		return nil
	},
	Run: func(cmd *Command, args []string) {

		c, ctx := newGravityRPCClient()
		defer c.close()
		ballHash:= removeMarks(ballHashString)
		res, err := c.GetBall(ctx, &pb.GetBallRequest{BallHash:ballHash})
		if err != nil {
			fmt.Println("Connection to the server failed")
			return
		}

		info,err:=json.MarshalIndent(res.Ball, "", " ")
		if err != nil {
			log.Errorf("Json marshal happened error，%v",err)
		}
		fmt.Printf("%+v \n", string(info))
	},
}

func init() {
	ballCmd.Flags().StringVarP(&ballHashString, "ball",
		"b","","ball [required]")
}

