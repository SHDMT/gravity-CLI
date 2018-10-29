package cmd

var getSeedCmd = &Command{
	Use:   "getseed",
	Short: "get seed for new wallet",
	Run: func(cmd *Command, args []string) {

		//c, ctx := newRpcClient()
		//defer func() {
		//	c.close()
		//	//退出命令行
		//	os.Exit(0)
		//}()

		//res, err := c.CreateSeed(ctx, &pb.CreateSeedRequest{})
		//
		//if err != nil {
		//	fmt.Print("=======result>>>>>>>")
		//	fmt.Print(err)
		//	fmt.Println("<<<<<<<result=======")
		//	return
		//}
		//fmt.Println(res.SeedMnemonic)
	},
}

