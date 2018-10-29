package cmd

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnMashalIssueMessage(t *testing.T){

	jsonString := `{
	"Header": {
		"App": 2,
		"Version": 1,
		"PayloadHash": ""
	},
	"Contracts": [{
		"Meta": {
			"Version": 1,
			"Name": "HelloWorld",
			"ScriptCode": 1,
			"IsRestrict": false,
			"ParamsKey": [
				"name1",
				"name2",
				"name3"
			],
			"ParamsNote": [
				"value1",
				"value2",
				"value3"
			]
		},
		"Code": "(Print HelloWorld)",
		"Note": "This is a HelloWorld Contract."
	}]
}`
	msg := UnMashalDeployMessage(jsonString)

	fmt.Printf("%+v \n", msg)
	payloadHash := msg.CalcPayloadHash()
	fmt.Printf("PayloadHash : %x \n ", payloadHash)

	data, _ := json.Marshal(msg)
	fmt.Println(string(data))
}