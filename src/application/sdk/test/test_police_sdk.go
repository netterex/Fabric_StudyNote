package main

import (
	"fabric-houserental/application/sdk/police"
	"fmt"
)

func main() {
	chaincode_name := "policeCc"
	fcn := "query"
	//args := [][]byte{[]byte("641111111111111111")}
	//resp, err := police.ChannelQuery(chaincode_name, fcn, args)
	//if err != nil {
	//	fmt.Println("query failed")
	//} else {
	//	fmt.Println("query success")
	//}
	//fmt.Println("===================")
	//fmt.Println(resp)

	fcn_exec := "set"
	set_args := [][]byte{[]byte("641111111111111112"), []byte("ls"), []byte("22"), []byte("yyy")}
	resp_set, err := police.ChannelExecute(chaincode_name, fcn_exec, set_args)
	
	if error.Error(err) == "Client Status Code: (5) TIMEOUT. Description: request timed out or been cancelled" {
		fmt.Println("set success")
	} else {
		fmt.Println("set failed")
	}
	
	fmt.Println(resp_set.ChaincodeStatus)

	verify_args := [][]byte{[]byte("641111111111111112")}
	resp_verify, err := police.ChannelQuery(chaincode_name, fcn, verify_args)
	if err != nil {
		fmt.Println("query failed")
	} else {
		fmt.Println("query success")
	}
	fmt.Println(resp_verify.ChaincodeStatus)
	fmt.Println(string(resp_verify.Payload))
}
