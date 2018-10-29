package cmd

import (
	"fmt"
	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
	"github.com/spf13/cobra"
	"github.com/pkg/errors"
	"os"
)
var unitHashString string
var unitCmd = &Command{
	Use:   "getunit",
	Short: "get unit by unitHash",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(unitHashString)==0 {
			return errors.New("unitHash has no input")
		}
		//TODO:
		//if len(unitHashString)!=44 {
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

		unitHash,err:= base64Parse(unitHashString)
		if err != nil {
			fmt.Println("Base64 conversion to binary error")
			return
		}
		res, err := c.GetUnitMessage(ctx,
			&pb.GetUnitMessageRequest{UnitHash:unitHash})
		if err != nil {
			fmt.Println("Connection to the server failed: ", err)
			return
		}

		//unit := new(structure.Unit)
		//buf := bytes.NewBuffer(res.Unit)
		//gob.NewDecoder(buf).Decode(unit)
		// unit := new(structure.Unit)
		// unit.Deserialize(res.Unit)
		// fmt.Printf("Bytes : %x \n", res.Unit)
		// fmt.Printf("Unit : %+v \n", unit)
		//info,err:=json.MarshalIndent(unit, "", " ")
		//if err != nil {
		//	log.Errorf("Json marshal happened error，%v",err)
		//}
		fmt.Println(string(res.Unit))
	},
}

func init() {
	unitCmd.Flags().StringVarP(&unitHashString, "unit",
		"u","","unit [required]")
}
