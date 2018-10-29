package cmd

var subCommands = []*Command{
	clearCmd,
	exitCmd,
	addressCmd,
	balanceCmd,
	historyCmd,
	messageCmd,
	textCmd,
	paymentsCmd,
	invokeContractCmd,
	issueContractCmd,
	deployContractCmd,
	sendrawunitCmd,
	getSeedCmd,
	validateAddressCmd,
	dumpCmd,
	importCmd,
	accountCmd,
	updatePasswordCmd,
	rescanWalletCmd,
	allAssetsCmd,
	invokeWithJsonCmd,
	issueWithJsonCmd,
	deployWithJsonCmd,
	//gravity
	startMiningCmd,
	stopMiningCmd,
	lastKeyUnitCmd,
	tipsCmd,
	keyUnitCmd,
	unitsCmd,
	infoCmd,
	isGoodCmd,
	unitBallCmd,
	ballCmd,
	isStableCmd,
	unitCmd,
	committeeCmd,
	currentBlockCmd,
	currentKeyMciCmd,
	blockCmd,
	getPeerInfoCmd,
	isCommitteeCmd,
	getMiningSpeedCmd,
	contractInfoCmd,
	//contractInfoByHashCmd,
	listAllContractsCmd,
	getNameOfAssetCmd,
	listContractsOfAssetCmd,
	progressCmd,
}

// the global unexported commander for CLI handling
var defaultCommander *commander

func init() {
	defaultCommander = newDefaultCommander()
	// register all commands specified
	defaultCommander.RegisterCommands(subCommands...)
	defaultCommander.RegisterCommands(helpCmd)

}

// Execute bootstraps the default Commander running in normal mode
func Execute() error {
	return defaultCommander.Execute()
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
//func Execute() {
//	if err := RootCmd.Execute(); err != nil {
//		fmt.Println(err)
//		os.Exit(-1)
//	}
//}

// ReadlineAndExecute bootstraps the default Commander running in ReadLine mode
func ReadlineAndExecute() error {

	return defaultCommander.ReadlineAndExecute()
}
