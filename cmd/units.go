package cmd
import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"github.com/spf13/cobra"
	"math"
	"github.com/pkg/errors"
	"encoding/json"
	"github.com/SHDMT/gravity/infrastructure/log"
	"os"
)
var unitsMci uint64
var unitsCmd = &Command{
	Use:   "getunits",
	Short: "get units from dag",
	Args: func(cmd *cobra.Command, args []string) error {
		if unitsMci==math.MaxUint64 {
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
		res, err := c.GetUnits(ctx, &pb.GetUnitsRequest{Mci:unitsMci})
		if err != nil {
			fmt.Println("Connection to the server failed")
			return
		}

		info,err:=json.MarshalIndent(res.UnitsHash, "", " ")
		if err != nil {
			log.Errorf("Json marshal happened error，%v",err)
		}

		fmt.Println(string(info))

	},
}

func init() {
	unitsCmd.Flags().Uint64VarP(&unitsMci, "mci", "m",
		math.MaxUint64,"mci's value [required]")
}