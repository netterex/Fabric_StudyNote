package main

import (
	"fabric-houserental/application/sdk/housemng"
	"fmt"
)

func main() {
	chaincode_name := "housemngCc"
	fcn := "query"
	args := [][]byte{[]byte("00000001")}
	resp, err := housemng.ChannelQuery(chaincode_name, fcn, args)
	if err != nil {
		fmt.Println("query failed")
	} else {
		fmt.Println("query success")
	}
	fmt.Println("===================")
	fmt.Println(resp)
	fmt.Println(resp.ChaincodeStatus)
	fmt.Println(string(resp.Payload))
}
