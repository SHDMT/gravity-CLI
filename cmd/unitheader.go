package cmd
//import (
//	"fmt"
//	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
//	"github.com/spf13/cobra"
//	"github.com/pkg/errors"
//	"encoding/json"
//	"github.com/SHDMT/gravity/infrastructure/log"
//	"os"
//)
//var unitHeaderString string
//var unitHeaderCmd = &Command{
//	Use:   "getunitheader",
//	Short: "get unit header by unitHash",
//	Args: func(cmd *cobra.Command, args []string) error {
//		if len(unitHeaderString)==0 {
//			return errors.New("Unit has no input")
//		}
//		//if len(unitHeaderString)!=44 {
//		//	return errors.New("input unitHash is wrong!")
//		//}
//		return nil
//	},
//	Run: func(cmd *Command, args []string) {
//
//		c, ctx := newRpcClient()
//		defer func() {
//			c.close()
//			//退出命令行
//			os.Exit(0)
//		}()
//
//		unitHash:=removeMarks(unitHeaderString)
//
//		res, err := c.GetUnitHeader(ctx,
//			&pb.GetUnitHeaderRequest{UnitHash:unitHash})
//		if err != nil {
//			fmt.Println("Connection to the server failed")
//			return
//		}
//
//		info,err:=json.MarshalIndent(res.Header, "", " ")
//		if err != nil {
//			log.Errorf("Json marshal happened error，%v",err)
//		}
//		fmt.Println(string(info))
//	},
//}
//
//func init() {
//	unitHeaderCmd.Flags().StringVarP(&unitHeaderString, "unit",
//		"u","","unit [required]")
//}