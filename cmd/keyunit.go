package cmd
import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"github.com/spf13/cobra"
	"math"
	"github.com/pkg/errors"
	"os"
)
var keyUnitMci uint64
var keyUnitCmd = &Command{
	Use:   "getkeyunit",
	Short: "get key unit from dag",
	Args: func(cmd *cobra.Command, args []string) error {
		if keyUnitMci==math.MaxUint64 {
			return errors.New("Mci has no input")
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

		resMci, err := c.GetCurrentKeyMci(ctx, &pb.CurrentKeyMciRequest{})
		if err != nil{
			fmt.Println("Connection to the server failed")
			return
		}

		if resMci.Mci < keyUnitMci {
			fmt.Println("The input MCI is bigger than the current maximum" +
				" MCI")
			return
		}
		res, err := c.GetKeyUnit(ctx, &pb.GetKeyUnitRequest{Mci:keyUnitMci})
		if err != nil {
			fmt.Println("Connection to the server failed")
			return
		}
		fmt.Println(res.KeyUnitHash)
	},
}

func init() {
	keyUnitCmd.Flags().Uint64VarP(&keyUnitMci, "mci", "m",
		math.MaxUint64,"mci's value [required]")
}

