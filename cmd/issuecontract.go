package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"encoding/hex"
	"encoding/base64"
	"fmt"

	pb "github.com/SHDMT/gwallet/platform/grpc/walletrpc"
	"errors"
	"strconv"
)


var assetCap		 	int64
var fixDenominations 	string
var denominations	  	[]string
var contractDefs		[]string
var allocationAddr		[]string
var allocationAmount	[]string
var publisherAddr 		string
var note 				string


// paymentCmd sends payment to the connected wallet
var issueContractCmd = &Command{
	Use:   "issuecontract",
	Short: "send issuecontract command to the connected wallet",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(contractDefs) == 0 {
			return errors.New("contract list must not empty")
		}
		return nil
	},
	Run: func(cmd *Command, args []string) {

		c, ctx := newGwalletRPCClient()
		defer func() {
			c.close()
			//退出命令行
			os.Exit(0)
		}()

		contracts :=make([][]byte, len(contractDefs))
		for i,contractStr := range contractDefs {
			fmt.Printf(" contract Def string : %s \n", contractStr)
			contract, err := hex.DecodeString(contractStr)
			if err != nil {
				fmt.Println(err)
				return
			}
			contracts[i] = contract
		}

		denominationsList := make([]uint32, len(denominations))
		for i,denominationStr := range denominations {
			denomination,err := strconv.ParseUint(denominationStr,10,64)
			fmt.Printf("denomination is : %d \n", denomination)
			if err != nil {
				fmt.Println(" your denomination is invalid")
				return
			}

			denominationsList[i] = uint32(denomination)
		}

		allocationAddrList := make([][]byte, len(allocationAddr))
		for i,addrStr := range allocationAddr {
			fmt.Printf(" allocationAddr is : %s \n ", addrStr)
			addrHash, err := base64.StdEncoding.DecodeString(addrStr)
			if err != nil {
				fmt.Println(" your allocationAddr is invalid")
				return
			}
			allocationAddrList[i] = addrHash
		}

		allocationAmountList := make([]int64, len(allocationAmount))
		for i,amountStr := range allocationAmount {
			amount, err := strconv.ParseInt(amountStr, 10,64)
			if err != nil {
				fmt.Println(" your allocationAmount is invalid")
				return
			}
			fmt.Printf(" allocationAddr is : %s \n ", amount)
			allocationAmountList[i] = amount
		}

		var publisherAddress []byte
		if publisherAddr != "" || len(publisherAddr) != 0{
			var err error
			fmt.Printf("publisherAddress is : %s \n",publisherAddr)
			publisherAddress, err = base64.StdEncoding.DecodeString(publisherAddr)
			if err != nil {
				fmt.Println(" your publisher address is invalid")
				return
			}
		}
		fixDenominationsBool := false
		if fixDenominations == "true"{
			 fixDenominationsBool = true
		}
		issueRequest := &pb.IssueContractRequest{
			AssetName:				assetName,
			AssetCap: 				assetCap,
			FixedDenominations:		fixDenominationsBool,
			Denominations: 			denominationsList,
			Contracts: 				contracts,
			AllocationAddr:			allocationAddrList,
			AllocationAmount:		allocationAmountList,
			PublisherAddress:		publisherAddress,
			Note: 					[]byte(note),
			Send:					send,
		}

		res, err := c.IssueContract(ctx, issueRequest)
		if err != nil {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(err)
			fmt.Println("<<<<<<<result=======")
			return
		}

		if !send {
			fmt.Print("=======result>>>>>>>")
			fmt.Print(res.Commission)
			fmt.Println("<<<<<<<result=======")
			return
		}
		fmt.Print("=======result>>>>>>>")
		fmt.Print(base64.StdEncoding.EncodeToString(res.UnitHash))
		fmt.Println("<<<<<<<result=======")
	},
}

func init() {
	issueContractCmd.Flags().StringVarP(&accountString, "accountName", "a", "",
		"account Name string [optional]")
	issueContractCmd.Flags().StringVarP(&assetName, "assetName", "b", "",
		"asset Name string [required]")
	issueContractCmd.Flags().Int64VarP(&assetCap, "assetCap", "c", 0,
		"asset cap uint64 [required]")
	issueContractCmd.Flags().StringVarP(&fixDenominations, "fixDenomination", "d", "",
		"fixDenomination [required]")
	issueContractCmd.Flags().StringSliceVarP(&denominations, "denominations", "e", nil,
		"denominations [optional]")
	issueContractCmd.Flags().StringSliceVarP(&contractDefs, "contractDefs", "f", nil,
		"contractDefs [required]")
	issueContractCmd.Flags().StringSliceVarP(&allocationAddr, "allocationAddr", "g", nil,
		"allocationAddr [required]")
	issueContractCmd.Flags().StringSliceVarP(&allocationAmount, "allocationAmount", "i", nil,
		"allocationAmount [required]")
	issueContractCmd.Flags().StringVarP(&publisherAddr, "publisherAddr", "j", "",
		"publisher Address  string [optional]")
	issueContractCmd.Flags().StringVarP(&note, "note", "k", "",
		"descript for the contract [optional]")
	issueContractCmd.Flags().BoolVarP(&send, "send", "s", true,
		"send the new unit to network or not [optional]")
}
