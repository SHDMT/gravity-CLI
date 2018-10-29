package cmd
//import (
//	"fmt"
//	pb "github.com/SHDMT/gravity/platform/grpc/gravityrpc"
//	"github.com/spf13/cobra"
//	"github.com/pkg/errors"
//	"encoding/base64"
//)
//var address string
//var definitionCmd = &Command{
//	Use:   "getdefinition",
//	Short: "get definition by address",
//	Args: func(cmd *cobra.Command, args []string) error {
//		if len(address)==0 {
//			return errors.New("Address has no input")
//		}
//		return nil
//	},
//	Run: func(cmd *Command, args []string) {
//
//		c, ctx := newRpcClient()
//		defer c.close()
//
//		addressHash,err:= base64Parse(address)
//		if err != nil {
//			fmt.Println("Base64 conversion to binary error")
//			return
//		}
//		res, err := c.GetDefinition(ctx,
//			&pb.GetDefinitionRequest{Address:addressHash})
//		if err != nil {
//			fmt.Println("Connection to the server failed")
//			return
//		}
//		fmt.Println("Definition: ",base64.StdEncoding.EncodeToString(res.
//			Definition))
//	},
//}
//
//func init() {
//	definitionCmd.Flags().StringVarP(&address, "address",
//		"a","","address [required]")
//}
