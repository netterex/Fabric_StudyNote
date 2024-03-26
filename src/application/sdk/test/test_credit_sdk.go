package main

import (
	"fabric-houserental/application/sdk/credit"
	"fmt"
)

func test_credit() {
	credit_chaincode_name := "creditCc"
	fcn_query := "query"
	//credit_args := [][]byte{[]byte("641111111111111111")}
	//credit_resp := credit.ChannelQuery(credit_chaincode_name, fcn_query, credit_args)
	//fmt.Println("==========creditCc=========")
	//fmt.Println(credit_resp)
	//fmt.Println(credit_resp.ChaincodeStatus)
	//fmt.Println(string(credit_resp.Payload))
	//
	fcn_exec := "set"
	set_args := [][]byte{[]byte("641111111111111112"), []byte("B")}
	resp_set, err := credit.ChannelExecute(credit_chaincode_name, fcn_exec, set_args)

	if error.Error(err) == "Client Status Code: (5) TIMEOUT. Description: request timed out or been cancelled" {
		fmt.Println("set success")
	} else {
		fmt.Println("set failed")
	}

	fmt.Println(resp_set.ChaincodeStatus)

	verify_args := [][]byte{[]byte("641111111111111112")}
	resp_verify, err := credit.ChannelQuery(credit_chaincode_name, fcn_query, verify_args)
	if err != nil {
		fmt.Println("query failed")
	} else {
		fmt.Println("query success")
	}
	fmt.Println(resp_verify.ChaincodeStatus)
	fmt.Println(string(resp_verify.Payload))

}

func test_contract() {
	contract_chaincode_name := "contractCc"
	fcn := "query"
	contract_args := [][]byte{[]byte("0000000001"), []byte("4")}
	contract_resp, err := credit.ChannelQuery(contract_chaincode_name, fcn, contract_args)
	if err != nil {
		fmt.Println("query failed")
	} else {
		fmt.Println("query success")
	}
	fmt.Println("==========contractCc=========")
	fmt.Println(contract_resp)
	fmt.Println(contract_resp.ChaincodeStatus)
	fmt.Println(string(contract_resp.Payload))
}

func test_transaction() {
	transaction_chaincode_name := "transactionCc"
	fcn := "query"
	transaction_args := [][]byte{[]byte("641111111111111111")}
	transaction_resp, err := credit.ChannelQuery(transaction_chaincode_name, fcn, transaction_args)
	if err != nil {
		fmt.Println("query failed")
	} else {
		fmt.Println("query success")
	}
	fmt.Println("==========transactionCc=========")
	fmt.Println(transaction_resp)
	fmt.Println(transaction_resp.ChaincodeStatus)
	fmt.Println(string(transaction_resp.Payload))
}

func test_user() {
	user_chaincode_name := "userCc"
	fcn := "query"
	user_args := [][]byte{[]byte("4-user-123456-18-female")}
	user_resp, err := credit.ChannelQuery(user_chaincode_name, fcn, user_args)
	if err != nil {
		fmt.Println("query failed")
	} else {
		fmt.Println("query success")
	}
	fmt.Println("==========userCc=========")
	fmt.Println(user_resp)
	fmt.Println(user_resp.ChaincodeStatus)
	fmt.Println(string(user_resp.Payload))
}

func main() {
	test_credit()

}
