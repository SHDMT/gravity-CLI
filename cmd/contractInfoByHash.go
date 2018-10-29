package cmd

var contractHash string
var contractInfoByHashCmd = &Command{
	//Use:   "getcontractinfo",
	//Short: "Get contract info",
	//Args: func(cmd *cobra.Command, args []string) error {
	//	if len(contractHash)==0 {
	//		return errors.New("contractHash has no input")
	//	}
	//	return nil
	//},
	//Run: func(cmd *Command, args []string) {
	//
	//	c, ctx := newGravityRPCClient()
	//	defer func() {
	//		c.close()
	//		//退出命令行
	//		os.Exit(0)
	//	}()
	//
	//	caddrHash,err:= base64Parse(contractHash)
	//	if err != nil {
	//		fmt.Println("Base64 conversion to binary error")
	//		return
	//	}
	//	contractInfoRequest:=&pb.ContractInfoRequest{Caddr:caddrHash}
	//	res, err := c.GetContractInfo(ctx,contractInfoRequest)
	//
	//	if err != nil {
	//		fmt.Print("=======result>>>>>>>")
	//		fmt.Print("Connection to the server failed")
	//		fmt.Println("<<<<<<<result=======")
	//		return
	//	}
	//
	//	info,err:=json.MarshalIndent(res.Contract, "", " ")
	//	if err != nil {
	//		log.Errorf("Json marshal happened error，%v",err)
	//	}
	//
	//	fmt.Print("=======result>>>>>>>")
	//	fmt.Print(string(info))
	//	fmt.Println("<<<<<<<result=======")
	//},
}

func init() {
	contractInfoByHashCmd.Flags().StringVarP(&contractHash, "contractHash",
		"c",""," contractHash [required]")
}