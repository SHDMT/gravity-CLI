package cmd

import (
	"encoding/hex"
	"fmt"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func Test_main(t *testing.T) {
	Execute()
	if err := ReadlineAndExecute(); nil != err {
		fmt.Println(err)
	}

}
func Test_HelpAndClear(t *testing.T) {

	RootCmd.Run(RootCmd, nil)
	helpCmd.Run(helpCmd, nil)
	clearCmd.Run(clearCmd, nil)

	resetAllFlagsToDefault(nil)
	resetAllFlagsToDefault(addressCmd)

}
func Test_Address(t *testing.T) {

	args := []string{"1", "2", "3", "4"}
	if addressCmd.Args(addressCmd, args) == nil {

		addressCmd.Run(addressCmd, args)

	} else {
		t.Fatalf("此处测试出错%v", ballotX)
	}

}
//func Test_Keyunit(t *testing.T) {
//
//	args := []string{"1", "2"}
//	if keyUnitCmd.Args(keyUnitCmd, args) == nil {
//
//		keyUnitCmd.Run(keyUnitCmd, args)
//
//	} else {
//		t.Fatalf("此处测试出错%v", ballotX)
//	}
//
//}

func Test_Balance(t *testing.T) {

	args := []string{}
	if balanceCmd.Args(balanceCmd, args) == nil {

		balanceCmd.Run(balanceCmd, args)

	} else {
		t.Fatalf("此处测试出错%v", ballotX)
	}
}

func Test_History(t *testing.T) {
	args := []string{"1", "2", "3", "4", "5", "6"}
	if historyCmd.Args(historyCmd, args) == nil {

		historyCmd.Run(historyCmd, args)

	} else {
		t.Fatalf("此处测试出错%v", ballotX)
	}
}

func Test_Message(t *testing.T) {
	unitHash := []byte{158, 98, 145, 151, 12, 180, 77, 217, 64, 8, 199, 155, 202,
		249, 216, 111, 24, 180, 180, 155, 165, 178, 160, 71, 129, 219, 113,
		153, 237, 59, 158, 78,
	}

	unitHashToString = hex.EncodeToString(unitHash)
	messageId = 4294967295
	args := []string{unitHashToString, string(messageId)}
	if messageCmd.Args(messageCmd, args) == nil {
		t.Fatalf("此处测试出错%v", ballotX)
	}

	unitHashToString = hex.EncodeToString(unitHash)
	messageId = 0
	args = []string{unitHashToString, string(messageId)}
	if messageCmd.Args(messageCmd, args) == nil {
		messageCmd.Run(messageCmd, args)
	} else {
		t.Fatalf("此处测试出错%v", ballotX)
	}

}

//func Test_Payment(t *testing.T) {
//	toAddress := []byte{158, 98, 145, 151, 12, 180, 77, 217, 64, 8, 199, 155, 202,
//		249, 216, 111, 24, 180, 180, 155, 165, 178, 160, 71, 129, 219, 113,
//		153, 237, 59, 158, 78,
//	}
//	//toAddressToString=hex.EncodeToString(toAddress)
//
//	//
//	//toAddressToString="123"
//	//amount=1
//	//args=[]string{toAddressToString,string(amount)}
//	//if paymentCmd.Args(paymentCmd, args) == nil {
//	//	t.Fatalf("此处测试出错%v", ballotX)
//	//}
//	toAddressToString = hex.EncodeToString(toAddress)
//	amount = 0
//	args := []string{}
//	if paymentCmd.Args(paymentCmd, args) == nil {
//		t.Fatalf("此处测试出错%v", ballotX)
//	}
//	toAddressToString = "3iLbp29LdZawygHSNo4p3v1s4ERQZZpkyNb3AbnJegCHFBqe4p"
//	amount = 1
//	args = []string{toAddressToString, string(amount)}
//	if paymentCmd.Args(paymentCmd, args) == nil {
//		paymentCmd.Run(paymentCmd, args)
//	} else {
//		t.Fatalf("此处测试出错%v", ballotX)
//	}
//}

func Test_Text(t *testing.T) {
	args := []string{"1", "2", "3", "4"}
	text = "123"
	if textCmd.Args(textCmd, args) == nil {
		textCmd.Run(textCmd, args)
	} else {
		t.Fatalf("此处测试出错%v", ballotX)
	}

	text = ""
	if textCmd.Args(textCmd, args) == nil {
		t.Fatalf("此处测试出错%v", ballotX)
	}
}

func Test_Exit(t *testing.T) {
	exitCmd.Run(exitCmd, nil)
}

func Test_notExistCmd(t *testing.T) {
	notExistCmd := &Command{}
	exitCmd.Run(notExistCmd, nil)

}

//func TestReadlineAndExecute(t *testing.T) {
//	args:=[]string{"1","2","3","4","5"}
//	if addressCmd.Args(addressCmd, args) == nil {
//
//		addressCmd.Run(addressCmd, args)
//
//	} else {
//		t.Fatalf("此处测试出错%v", ballotX)
//	}
//
//	instance := &readline.Instance{
//		Config  :  nil,
//		Terminal :  nil,
//		Operation:  nil,
//	}
//	notExistCmd :=&commander{
//		rootCmd:RootCmd,           // the command at the highest level
//		subCmds: nil,// collection of all 2nd-level commands
//		rl:instance,
//	}
//	err:=notExistCmd.ReadlineAndExecute()
//	if err != nil {
//		t.Logf("测试通过%v",checkMark)
//	}
//	resetAllFlagsToDefault(nil)
//	resetAllFlagsToDefault(addressCmd)
//}
