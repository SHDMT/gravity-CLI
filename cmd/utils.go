package cmd

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/SHDMT/gravity/infrastructure/log"
	"github.com/SHDMT/gravity/platform/consensus/structure"
	"github.com/pkg/errors"
)

//去除命令行输入的双引号
func removeMarks(stringValue string) string {
	stringValueBytes := []byte(stringValue)
	stringValueBytesLen := len(stringValueBytes)
	if stringValueBytesLen < 1 {
		return stringValue
	}
	if stringValueBytes[0] == 34 &&
		stringValueBytes[stringValueBytesLen-1] == 34 {
		noMarksStringBytes := stringValueBytes[1 : stringValueBytesLen-1]
		stringValueBytes = noMarksStringBytes
	}
	noMarksString := string(stringValueBytes)
	return noMarksString
}

//base64转换为二进制
func base64Parse(h string) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(removeMarks(h))
	if err != nil {
		return nil,errors.Wrap(err,"Base64 conversion to binary error")
	}
	return b,nil
}

//UnMashalIssueMessage
func UnMashalIssueMessage(issueJson string)(*structure.IssueMessage){
	issueBytes:=[]byte(issueJson)
	issueMessage := new(structure.IssueMessage)
	err := json.Unmarshal(issueBytes, issueMessage)
	if err != nil{
		log.Error("failed to Unmashal")
	}
	return issueMessage
}

//UnMashalIssueMessage
func UnMashalInvokeMessage(invokeJson string)(*structure.InvokeMessage){
	invokeBytes:=[]byte(invokeJson)
	invokeMessage := new(structure.InvokeMessage)
	err := json.Unmarshal(invokeBytes, invokeMessage)
	if err != nil{
		log.Error("failed to Unmashal")
	}
	return invokeMessage
}

//UnMashalDeployMessage
func UnMashalDeployMessage(deployJson string)(*structure.DeployMessage){
	//deployBytes:=[]byte(deployJson)
	data ,_:=hex.DecodeString("7b0a0922486561646572223a207b0a090922417070223a20322c0a09092256657273696f6e223a20312c0a0909225061796c6f616448617368223a2022220a097d2c0a0922436f6e747261637473223a205b7b0a0909224d657461223a207b0a0909092256657273696f6e223a20312c0a090909224e616d65223a202248656c6c6f576f726c64222c0a09090922536372697074436f6465223a20312c0a0909092249735265737472696374223a2066616c73652c0a09090922506172616d734b6579223a205b0a09090909226e616d6531222c0a09090909226e616d6532222c0a09090909226e616d6533220a0909095d2c0a09090922506172616d734e6f7465223a205b0a090909092276616c756531222c0a090909092276616c756532222c0a090909092276616c756533220a0909095d0a09097d2c0a090922436f6465223a2022285072696e742048656c6c6f576f726c6429222c0a0909224e6f7465223a20225468697320697320612048656c6c6f576f726c6420436f6e74726163742e220a097d5d0a7d")
	fmt.Printf("%s \n", data)
	deployMessage := new(structure.DeployMessage)
	err := json.Unmarshal(data, deployMessage)
	if err != nil{
		log.Error("failed to Unmarshal")
	}
	return deployMessage
}